package main

import (
	"log"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/handler"
	"github.com/vignesh-j-shetty/GoDFS/internal/mainnode/service"
)

func main() {
	// Load env variables
	_ = godotenv.Load()
	config := config.InitConfig()
	ZookeeperService, err := service.NewZookeeperService(config)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Gin setup
	r := gin.Default()
	metadataService := service.NewMetaDataService(config, ZookeeperService)
	folderOps := handler.NewFolderOpsHandler(&metadataService)
	folderOps.InitRoutes(r)

	fileOps := handler.NewFileOpsHandler(&metadataService)
	fileOps.InitRoutes(r)
	
	var wg sync.WaitGroup
	wg.Add(2)

	go func ()  {
		defer wg.Done()
		ZookeeperService.WatchLoop()
	}()

	go func() {
		defer wg.Done()
		log.Printf("https server listening on 8080")
		r.Run(":8080")
	}()

	wg.Wait()
}