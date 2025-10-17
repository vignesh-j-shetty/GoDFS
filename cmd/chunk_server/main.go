package main

import (
	"fmt"

	"github.com/vignesh-j-shetty/GoDFS/internal/chunkserver"
)

func main() {
	chunk_client, err := chunkserver.NewServerConnector()

	if err != nil {
		fmt.Printf("Failed to connect : %s", err)
	}

	chunk_client.ConnectWithServer()
}