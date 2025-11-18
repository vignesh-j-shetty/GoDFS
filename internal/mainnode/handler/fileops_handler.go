package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	restapi "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api"
)

type FileOpsHandler struct {
	MetaDataHandler *service.MetaDataService
	zookeeperService *service.ZookeeperService
}

func NewFileOpsHandler(metaDataService *service.MetaDataService, zookeeperService *service.ZookeeperService) FileOpsHandler {
	return FileOpsHandler{
		MetaDataHandler: metaDataService,
		zookeeperService: zookeeperService,
	}
}

func (f *FileOpsHandler) InitRoutes(r *gin.Engine) {
	r.POST("/v1/file/create", f.createNewFile)
	r.POST("/v1/file/chunkinfo", f.getFileChunkInfo)
}

func (f *FileOpsHandler) createNewFile(ctx *gin.Context) {
	var createNewFile restapi.FileCreateRequest

	if !bindOrAbort(ctx, &createNewFile) {
		return
	}

	chunkLocations, err := f.MetaDataHandler.CreateFiles(createNewFile.Path, createNewFile.FileName, createNewFile.Size)

	if err != nil {
		resp := restapi.Response{
			Status: commonconstants.FAILURE_STATUS,
			Error:  err.Error(),
			Data:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	var chunkUploadInfoList []restapi.ChunkInfo

	for _, chunkLocation := range chunkLocations {
		var uploadUrls []string
		for _, DataNodeID := range chunkLocation.DataNodeIDs {
			uploadUrls = append(uploadUrls, DataNodeID.UploadUrl)
		}
		chunkUploadInfoList = append(chunkUploadInfoList, restapi.ChunkInfo{
			ChunkId:      chunkLocation.ChunkID,
			DataNodeUrls: uploadUrls,
		})
	}

	resp := restapi.Response{
		Status: commonconstants.SUCCESS_STATUS,
		Data:   chunkUploadInfoList,
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (f *FileOpsHandler) getFileChunkInfo(ctx *gin.Context) {
	var fileChunkInfoRequest restapi.GetFileInfoRequest
	if !bindOrAbort(ctx, &fileChunkInfoRequest) {
		return
	}

	chunkInfo, err := f.MetaDataHandler.GetFileChunkList(fileChunkInfoRequest.Path)
	if err != nil {
		resp := restapi.Response{
			Status: commonconstants.FAILURE_STATUS,
			Error:  err.Error(),
			Data:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	var dataNodeInfo []restapi.ChunkInfo

	for _, chunk := range chunkInfo {
		chunkLocationInfo := f.zookeeperService.GetDataNodeForChunk(chunk)
		if len(chunkLocationInfo) == 0 {
			dataNodeInfo = append(dataNodeInfo, restapi.ChunkInfo{
				ChunkId: "Null",
			})
		} else {
			dataNodeInfo = append(dataNodeInfo, restapi.ChunkInfo{
				ChunkId: chunkLocationInfo[0].Id,
			})
		}

	}
	resp := restapi.Response{
		Status: commonconstants.SUCCESS_STATUS,
		Data:   dataNodeInfo,
	}
	ctx.JSON(http.StatusAccepted, resp)
}
