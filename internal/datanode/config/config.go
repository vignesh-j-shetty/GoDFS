package config

import (
	"os"
	"strings"
	"github.com/google/uuid"
)

type Config struct {
	MainNodeRCPUrl string
	UploadUrl string
	ChunkFilePath string
	ZookeepersServers []string
	Id string
}

func InitConfig() Config {
	randomUnqiueID := uuid.New()
	return Config{
		MainNodeRCPUrl: os.Getenv("MAINNODE_RPC_URL"),
		UploadUrl: os.Getenv("UPLOAD_URL"),
		ChunkFilePath: os.Getenv("CHUNK_FILE_PATH"),
		ZookeepersServers: strings.Split(os.Getenv("ZOOKEEPER_SERVERS"), ","),
		Id: randomUnqiueID.String(),
	}
}