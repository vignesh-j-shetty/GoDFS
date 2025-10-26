package config

import "os"

type Config struct {
	MainNodeRCPUrl string
}

func InitConfig() Config {
	return Config{
		MainNodeRCPUrl: os.Getenv("MAINNODE_RPC_URL"),
	}
}