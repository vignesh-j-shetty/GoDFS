package config

import "os"

type Config struct {
	MainNodeRCPUrl string
	SelfUrl string
	ChunkFilePath string
}

func InitConfig() Config {
	return Config{
		MainNodeRCPUrl: os.Getenv("MAINNODE_RPC_URL"),
		SelfUrl: os.Getenv("SELF_URL"),
		ChunkFilePath: os.Getenv("CHUNK_FILE_PATH"),
	}
}