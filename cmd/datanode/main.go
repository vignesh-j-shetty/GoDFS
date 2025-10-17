package main

import (
	"fmt"
	"github.com/vignesh-j-shetty/GoDFS/internal/datanode"
)

func main() {
	chunk_client, err := datanode.NewServerConnector()

	if err != nil {
		fmt.Printf("Failed to connect : %s", err)
	}

	chunk_client.ConnectWithServer()
}