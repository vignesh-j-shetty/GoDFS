package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
)

type FileUploadHandler struct {
	Config config.Config
}

func (f *FileUploadHandler) InitRoutes(r *gin.Engine) {
	r.POST("/v1/file/upload", f.upload)
}

func (f *FileUploadHandler) upload(c *gin.Context) {
	
	chunkFile, err := c.FormFile("chunk_file")
	if err != nil {
        c.String(http.StatusBadRequest, "File missing: %v", err)
        return
    }
	savePath := filepath.Join(f.Config.ChunkFilePath, chunkFile.Filename)
	err = c.SaveUploadedFile(chunkFile, savePath)

	if err != nil {
		c.String(http.StatusInternalServerError, "Unexpected error")
		return
	}

	chunkID := c.PostForm("chunk_id")
	c.String(http.StatusOK, "File ID :%s File size :%d", chunkID, chunkFile.Size)
	c.String(http.StatusOK, " File name :%s", chunkFile.Filename)
}