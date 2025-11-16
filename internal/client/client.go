package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	restapi "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api"
)


type GoDFSClient struct {
	config *Config
	httpClient *http.Client
}

func NewGoDFSClient(config *Config) *GoDFSClient {
	return &GoDFSClient{
		config: config,
		httpClient: &http.Client{ 
            Timeout: 30 * time.Second,
        },
	}
}

func (c *GoDFSClient) CreateFile(uploadPath string, filePath string) error {
	filename := filepath.Base(filePath)
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fileRequest := restapi.FileCreateRequest {
		Path: uploadPath,
		FileName: filename,
		Size: uint64(info.Size()),
	}
	// Gets chunk info from metadata server
	chunkInfoList, err := c.createFileInMetaDataServer(fileRequest)
	if err != nil {
		fmt.Println("Error getting chunk info:", err)
		return err
	}

	// Read first ChunkSize bytes from file and upload to datanodes
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var outErrs []string

	for _, chunkInfo := range chunkInfoList {
		chunkData := make([]byte, c.config.ChunkSize)
		n, err := file.Read(chunkData)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading chunk data:", err)
			return err
		}
		if n == 0 {
			break
		}
		err = c.uploadChunkToAllReplicas(chunkInfo, chunkData)
		if err != nil {
			outErrs = append(outErrs, fmt.Sprintf("chunk %s: %v", chunkInfo.ChunkId, err))
		}
	}
	if len(outErrs) > 0 {
		return fmt.Errorf("one or more chunks failed to upload: %s", strings.Join(outErrs, "; "))
	}
	return nil
}

func (c *GoDFSClient) createFileInMetaDataServer(fileRequest restapi.FileCreateRequest) ([]restapi.ChunkInfo, error) {
	jsonData, err := json.Marshal(fileRequest)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	resp, err := http.Post(c.config.MetadataServer + "/v1/file/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var response restapi.Response

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	if response.Status != commonconstants.SUCCESS_STATUS {
		return nil, fmt.Errorf("metadata server error: %s", response.Error)
	}

	var chunkInfoList []restapi.ChunkInfo
	if dataBytes, err := json.Marshal(response.Data); err == nil {
		if err := json.Unmarshal(dataBytes, &chunkInfoList); err != nil {
			fmt.Println("Error unmarshaling chunk info:", err)
			return nil, err
		}
	}

	return chunkInfoList, nil
}

func (c *GoDFSClient) uploadChunkToAllReplicas(chunkInfo restapi.ChunkInfo, chunkData []byte) error {
	var wg sync.WaitGroup
	errs := make(chan error, len(chunkInfo.UploadUrl))

	for _, uploadUrl := range chunkInfo.UploadUrl {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := c.uploadChunk(url, chunkInfo.ChunkId, chunkData)
			if err != nil {
				errs <- err
			}
		}(uploadUrl)
	}

	wg.Wait()
	close(errs)
	
	var outErrs []string
    for e := range errs {
        outErrs = append(outErrs, e.Error())
    }
	if len(outErrs) > 0 {
        return fmt.Errorf("one or more uploads failed: %s", strings.Join(outErrs, "; "))
    }
	return nil
}

func (c *GoDFSClient) uploadChunk(url string, filename string, chunkData []byte) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("chunk_file", filename)

	if err != nil {
		fmt.Println("Error creating form file:", err)
		return err
	}

	if _, err := io.Copy(part, bytes.NewReader(chunkData)); err != nil {
		return fmt.Errorf("write chunk data to form file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("close multipart writer: %w", err)
	}

	req, err := http.NewRequest("POST", url + "/v1/chunk/upload", &b)
	if err != nil {
		return fmt.Errorf("create upload request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("perform upload request: %w", err)
	}
	defer resp.Body.Close()

	var response restapi.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return err
	}
	if response.Status != commonconstants.SUCCESS_STATUS {
		return fmt.Errorf("upload failed: %s", response.Error)
	}

	return nil
}