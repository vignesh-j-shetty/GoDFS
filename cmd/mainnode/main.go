package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/handler"
)

func main() {
	_ = godotenv.Load()
	r := gin.Default()

	
	folderOps := handler.NewFolderOpsHandler()
	folderOps.InitRoutes(r)

	r.Run(":8080")
}