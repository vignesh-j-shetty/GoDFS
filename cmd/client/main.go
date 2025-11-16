package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vignesh-j-shetty/GoDFS/internal/client"
)

func main() {
	_ = godotenv.Load()
	conf := client.LoadConfig()
	godfsClient := client.NewGoDFSClient(&conf)
	
	// Require at least 3 arguments: <command> <remotePath> <localPath>
	if len(os.Args) < 4 {
		fmt.Println("Usage:")
		fmt.Println("  godfs create <remotePath> <localFilePath>")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		remotePath := os.Args[2]
		localPath := os.Args[3]

		err := godfsClient.CreateFile(remotePath, localPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
		} else {
			fmt.Println("File created successfully")
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Supported commands: create")
	}
}