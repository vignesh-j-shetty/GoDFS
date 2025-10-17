package main

import (
	"fmt"

	"github.com/vignesh-j-shetty/GoDFS/internal/datanode"
)

func main() {
	connection, err := datanode.NewMainNodeConnector()

	if err != nil {
		fmt.Printf("Failed to connect : %s", err)
	}

	connection.ConnectWithServer()
}
