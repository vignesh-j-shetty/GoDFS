package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
)

func main() {
	_ = godotenv.Load()
	config := config.InitConfig()
	connection, err := datanode.NewMainNodeConnector(config)

	if err != nil {
		fmt.Printf("Failed to connect : %s", err)
	}
	err = connection.ConnectWithServer()
	if err != nil {
		fmt.Println(err.Error())
	}
}
