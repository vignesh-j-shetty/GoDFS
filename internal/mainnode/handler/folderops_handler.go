package handler

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
	restapimodels "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api"
)

type FolderOpsHandler struct {
	MetaDataHandler* service.MetaDataService
}

func NewFolderOpsHandler(metaDataService* service.MetaDataService) FolderOpsHandler {
	return FolderOpsHandler{MetaDataHandler: metaDataService}
}

func (f *FolderOpsHandler) InitRoutes(r *gin.Engine) {
	r.POST("/v1/folder/create", f.createFolder)
	r.DELETE("/v1/folder/delete/*path", f.deleteFolder)
	r.GET("/v1/folder/content/*path", f.getFolderContent)
}

func (f *FolderOpsHandler) createFolder(ctx *gin.Context) {
	var createfolder restapimodels.FolderRequest
	if !bindOrAbort(ctx, &createfolder) {
		return
	}

	if !f.isValidPath(createfolder.Path) || strings.Contains(createfolder.FolderName, ".") {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"FAILURE", "errorMsg": "path/foldername invalid"})
		return
	}

	err := f.MetaDataHandler.CreateFolder(createfolder.Path, createfolder.FolderName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"FAILURE","errorMsg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func (f *FolderOpsHandler) deleteFolder(ctx *gin.Context) {
    path := ctx.Param("path")

	if path == "" {
		ctx.JSON(http.StatusOK, gin.H{"status":"FAILURE", "errorMsg": "Bad request"})
		return
	}

	err := f.MetaDataHandler.Delete(path)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"FAILURE","errorMsg": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func (f *FolderOpsHandler) getFolderContent(ctx *gin.Context) {
	path := ctx.Param("path")

	if path == "" {
		ctx.JSON(http.StatusOK, gin.H{"status":"FAILURE","errorMsg": "Bad request"})
		return
	}
	content, err := f.MetaDataHandler.GetFolderContents(path)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status":"FAILURE", "errorMsg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restapimodels.FolderContentList {
		FolderContentList: content,
	})
}

func (f *FolderOpsHandler) isValidPath(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '/' {
			continue
		}
		return false
	}
	return true
}