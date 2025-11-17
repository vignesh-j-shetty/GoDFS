package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/handler"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/service"
)

func main() {
	_ = godotenv.Load()
	config := config.InitConfig()

	r := gin.Default()

	zookeeperService, err := service.NewZookeeperClientService(config)

	handler := handler.FileUploadHandler {
		Config: config,
		ZookeeperClientService: zookeeperService,
	}

	if err != nil {
		log.Fatalln("Error ", err.Error())
	}
	fmt.Println("Zoo keeper register calling")
	zookeeperService.Register()
	
	handler.InitRoutes(r)
	r.Run(":8080")
}
