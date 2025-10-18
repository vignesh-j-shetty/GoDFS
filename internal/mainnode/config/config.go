package config

import (
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	DataNodeCount int
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
	return Config {
		DatabaseURL: os.Getenv("DATABASE_URL"),
		DataNodeCount: dataNodeCount,
	}
}