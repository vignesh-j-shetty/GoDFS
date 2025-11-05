package config

import (
	"os"
	"strconv"
	"strings"
	"github.com/google/uuid"
)

type Config struct {
	MainNodeRCPUrl string
	UploadUrl string
	ChunkFilePath string
	ZookeepersServers []string
	Id string
	TotalCapacity uint64
}

func InitConfig() Config {
	randomUnqiueID := uuid.New()
	totalCapacityStr := os.Getenv("DATA_CAPACITY")
	totalCapacity, err := strconv.ParseUint(totalCapacityStr, 10, 64)
	if err != nil {
		totalCapacity = 1024
	}

	// Convert MB to bytes
	totalCapacity *= 1024 * 1024
	return Config{
		MainNodeRCPUrl: os.Getenv("MAINNODE_RPC_URL"),
		UploadUrl: os.Getenv("UPLOAD_URL"),
		ChunkFilePath: os.Getenv("CHUNK_FILE_PATH"),
		ZookeepersServers: strings.Split(os.Getenv("ZOOKEEPER_SERVERS"), ","),
		Id: randomUnqiueID.String(),
		TotalCapacity: totalCapacity,
	}
}