package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	commonconstants "github.com/vignesh-j-shetty/GoDFS/pkg/common-constants"
	restapi "github.com/vignesh-j-shetty/GoDFS/pkg/rest-api"
)


func bindOrAbort(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindBodyWithJSON(obj); err != nil {
		ctx.JSON(http.StatusBadRequest, restapi.Response {
			Status: commonconstants.FAILURE_STATUS,
			Error: "Bad Request, Invalid Json Format",
		})
		return false
	}
	return true
}