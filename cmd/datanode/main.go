package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/v3/disk"
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
	usage, err := disk.Usage("/data")
	if err != nil {
		log.Fatalf("Error getting disk usage for %s: %v", "/", err)
	}

	// The Free field gives you the total unallocated space on the filesystem.
	// This includes space that might be reserved for the root user.
	freeBytes := usage.Free

	totalBytes := usage.Total

	// Convert bytes to a more readable format (GB)
	const GB = 1024 * 1024 * 1024

	fmt.Printf("Total Space:    %d GB\n", totalBytes/GB)
	fmt.Printf("Free Space:     %d GB\n", freeBytes/GB)
	fmt.Printf("Used Percentage: %.2f%%\n", usage.UsedPercent)
}
