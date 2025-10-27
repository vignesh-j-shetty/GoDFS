package main

import (
	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
)

func main() {
	_ = godotenv.Load()
	config.InitConfig()
}
