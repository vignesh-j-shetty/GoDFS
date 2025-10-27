package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/handler"
)

func main() {
	_ = godotenv.Load()
	config := config.InitConfig()

	r := gin.Default()

	handler := handler.FileUploadHandler {
		Config: config,
	}

	handler.InitRoutes(r)
	r.Run(":8080")
}
