package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
	"github.com/vignesh-j-shetty/GoDFS/pkg/api"
)

type FolderOpsHandler struct {
	MetaDataHandler service.MetaDataService
}

func NewFolderOpsHandler() FolderOpsHandler {
	metaDataService := service.NewMetaDataService()
	return FolderOpsHandler{MetaDataHandler: metaDataService}
}

func (f *FolderOpsHandler) InitRoutes(r *gin.Engine) {
	r.POST("/v1/createfolder", f.createFolder)
	r.DELETE("/v1/delete/*path", f.deleteFolder)
}

func (f *FolderOpsHandler) createFolder(ctx *gin.Context) {
	var createfolder api.FolderRequest
	if !f.bindOrAbort(ctx, &createfolder) {
		return
	}
	err := f.MetaDataHandler.CreateFolder(createfolder.Path, createfolder.FolderName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func (f *FolderOpsHandler) deleteFolder(ctx *gin.Context) {
    path := ctx.Param("path")

	if path == "" {
		ctx.JSON(http.StatusOK, gin.H{"status": "Bad request"})
		return
	}

	err := f.MetaDataHandler.Delete(path)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func (f *FolderOpsHandler) bindOrAbort(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindBodyWithJSON(obj); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "Bad request"})
		return false
	}
	return true
}
