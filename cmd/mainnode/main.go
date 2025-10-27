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
		log.Fatalf(err.Error())
	}

	ZookeeperService.WatchLoop()
	// Gin setup
	r := gin.Default()
	folderOps := handler.NewFolderOpsHandler()
	folderOps.InitRoutes(r)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("https server listening on 8080")
		r.Run(":8080")
	}()

	wg.Wait()
}