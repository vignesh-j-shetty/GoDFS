package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DatabaseURL string
	DataNodeCount int
	ZookeepersServers []string
	ChunkSize uint64
}

func InitConfig() Config {
	countStr := os.Getenv("DATANODE_COUNT")
	var dataNodeCount int = 3
	if countStr != "" {
		_dataNodeCount, err := strconv.Atoi(countStr)
		if err == nil {
			dataNodeCount = _dataNodeCount
		}
	}
	
	chunkSize , err := strconv.ParseUint(os.Getenv("CHUNK_SIZE"), 10, 64)
	if err != nil {
		chunkSize = 100
	}
	// Convert from MB to bytes
	chunkSize *= 1024 * 1024

	return Config {
		DatabaseURL: os.Getenv("DATABASE_URL"),
		DataNodeCount: dataNodeCount,
		ZookeepersServers: strings.Split(os.Getenv("ZOOKEEPER_SERVERS"), ","),
		ChunkSize: chunkSize,
	}
}