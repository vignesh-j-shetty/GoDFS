package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/service"
)

func main() {
	_ = godotenv.Load()
	heartHeatbeatService := service.HeartbeatClientService {}
	config := config.InitConfig()
	heartHeatbeatService.StartHeartbeatService(config, context.Background())
	heartHeatbeatService.Wait()
}
