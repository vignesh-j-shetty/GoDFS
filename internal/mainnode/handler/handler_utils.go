package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func bindOrAbort(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindBodyWithJSON(obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
		return false
	}
	return true
}