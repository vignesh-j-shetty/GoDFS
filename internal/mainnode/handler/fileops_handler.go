package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
	restapimodels "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api-models"
)

type FileOpsHandler struct {
	MetaDataHandler service.MetaDataService
}

func NewFileOpsHandler() FileOpsHandler {
	return FileOpsHandler{
		MetaDataHandler: service.NewMetaDataService(),
	}
}

func (f *FileOpsHandler) InitRoutes(r *gin.Engine) {
	r.POST("/v1/file/create", f.createNewFile)
}

func (f *FileOpsHandler) createNewFile(ctx *gin.Context) {
	var createNewFile restapimodels.FileCreateRequest

	if !bindOrAbort(ctx, &createNewFile) {
		return
	}
	
}