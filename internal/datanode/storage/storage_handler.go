package storage

import (
	"fmt"
	"os"

	"github.com/vignesh-j-shetty/GoDFS/internal/datanode/config"
)

type StorageHandler struct {
	config config.Config
}

func NewStorageHandler(config config.Config) (StorageHandler) {
	return StorageHandler{
		config: config,
	}
}

func (storageHandler *StorageHandler) GetFreeSpace() (uint64, error) {
	chunkFilePath := storageHandler.config.ChunkFilePath

	files, err := os.ReadDir(chunkFilePath)

	if err != nil {
		return 0, err
	}
	var totalUsedSpace uint64 = 0
	for _, file := range files {
		info, err := file.Info()

		if err != nil {
			fmt.Printf("Failed to get the info of %s\n", file.Name())
			continue
		}

		totalUsedSpace += uint64(info.Size())
	}

	return storageHandler.config.TotalCapacity - totalUsedSpace, nil
}