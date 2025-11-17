package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/storage"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/model"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
)

type ZookeeperClientService struct {
	conn           *zk.Conn
	evCh           <-chan zk.Event
	config         config.Config
	storageHandler storage.StorageHandler
}

func NewZookeeperClientService(config config.Config) (*ZookeeperClientService, error) {
	conn, evCh, err := zk.Connect(config.ZookeepersServers, time.Second)

	if err != nil {
		return nil, err
	}

	return &ZookeeperClientService{
		conn:           conn,
		evCh:           evCh,
		config:         config,
		storageHandler: storage.NewStorageHandler(config),
	}, nil
}

func (s *ZookeeperClientService) Register() error {
	freeSpace, err := s.storageHandler.GetFreeSpace()

	if err != nil {
		fmt.Println("Failed to get disk info ", err.Error(), " for path ", s.config.ChunkFilePath)
		return err
	}

	fmt.Println("Disk size :", freeSpace)

	chunkInfo := model.DataNodeInfo{
		Id:        s.config.Id,
		UploadUrl: s.config.UploadUrl,
		FreeSpace: freeSpace,
	}
	jsonData, err := json.Marshal(chunkInfo)

	if err != nil {
		return err
	}

	_, err = s.conn.Create(commonconstants.ChunkServerPrefixPath+"/"+chunkInfo.Id, jsonData, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Println("Failed to register chunkserver with zookeeper ", err.Error())
		return err
	}

	_, err = s.conn.Create(commonconstants.ChunkFilePrefixPath + "/" + chunkInfo.Id, []byte{}, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Println("Failed to ensure chunk file path ", err.Error())
		return err
	}

	fmt.Printf("client successfully registered with zookeeper with id %s\n", chunkInfo.Id)
	return nil
}

func (s *ZookeeperClientService) UpdateChunks() error {
	files, err := os.ReadDir(s.config.ChunkFilePath)

	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return err
	}
	var filesNames []string
	for _, file := range files {
		filesNames = append(filesNames, file.Name())
	}
	allFileNames := strings.Join(filesNames, ",")

	path := commonconstants.ChunkFilePrefixPath + "/" + s.config.Id
	_, stat, err := s.conn.Get(path)

    if err != nil {
        return fmt.Errorf("failed to get znode: %w", err)
    }
	_, err = s.conn.Set(path, []byte(allFileNames), stat.Version)

    if err != nil {
        return fmt.Errorf("failed to update znode: %w", err)
    }

	fmt.Println("Chunk files updated in zookeeper:", allFileNames)
	return nil
}