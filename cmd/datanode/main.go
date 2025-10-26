package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/service"
)

func main() {
	_ = godotenv.Load()
	config := config.InitConfig()
	heartHeatbeatService := service.HeartbeatClientService {
		DataNodeConfig: config,
	}
	heartHeatbeatService.StartHeartbeatService(context.Background())
	heartHeatbeatService.Wait()
}
