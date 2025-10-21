package datanode

import (
	"log"
	"os"
	"github.com/google/uuid"
)

func GetNodeId(rootPath string) string {

	path := rootPath + string(os.PathSeparator) + "iNode"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		
		if err := os.MkdirAll(rootPath, 0755); err != nil {
            log.Fatalf("Failed to create folder: %v", err)
        }

		id := uuid.New().String()
		if err := os.WriteFile(path, []byte(id), 0644); err != nil {
            log.Fatalf("Failed to write file: %v", err)
        }

		return id
	}

	content, err := os.ReadFile(path)
	if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }
	return string(content)
}