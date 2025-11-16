package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	restapi "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api"
)

type FileUploadHandler struct {
	Config config.Config
}

func (f *FileUploadHandler) InitRoutes(r *gin.Engine) {
	r.POST(commonconstants.UPLOAD_FILE_ENDPOINT, f.upload)
}

func (f *FileUploadHandler) upload(c *gin.Context) {
	chunkFile, err := c.FormFile("chunk_file")
	if err != nil {
		errorResp := restapi.Response {
			Status: commonconstants.FAILURE_STATUS,
			Error:  "Invalid chunk file",
		}
		c.JSON(http.StatusBadRequest, errorResp)
        return
    }
	savePath := filepath.Join(f.Config.ChunkFilePath, chunkFile.Filename)
	err = c.SaveUploadedFile(chunkFile, savePath)

	if err != nil {
		errorResp := restapi.Response {
			Status: commonconstants.FAILURE_STATUS,
			Error:  "Failed to save chunk file",
		}
		c.JSON(http.StatusInternalServerError, errorResp)
		return
	}

	resp := restapi.Response {
		Status: commonconstants.SUCCESS_STATUS,
	}
	c.JSON(http.StatusOK, resp)
}