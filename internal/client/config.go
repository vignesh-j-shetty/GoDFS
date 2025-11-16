package client

import (
	"os"
	"strconv"
)

type Config struct {
	MetadataServer     string
	DefaultConcurrency int
	ChunkSize		 int64
}

func LoadConfig() Config {
	chunkSize , err := strconv.ParseUint(os.Getenv("CHUNK_SIZE"), 10, 64)
	if err != nil {
		chunkSize = 100
	}
	// Convert from MB to bytes
	chunkSize *= 1024 * 1024

	defaultConcurrency, err := strconv.Atoi(os.Getenv("DEFAULT_CONCURRENCY"))
	if err != nil {
		defaultConcurrency = 4
	}

	return Config{
		MetadataServer:     os.Getenv("METADATA_SERVER"),
		DefaultConcurrency: defaultConcurrency,
		ChunkSize: int64(chunkSize),
	}
}

