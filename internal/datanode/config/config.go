package config

import "os"

type Config struct {
	MainNodeUrl string
}

func InitConfig() Config {
	return Config{
		MainNodeUrl: os.Getenv("MAINNODE_URL"),
	}
}