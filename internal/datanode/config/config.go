package config

import "os"

type Config struct {
	MainNodeUrl string
	SelfUrl string
	ListenPort string
	RootDataFolder string
}

func InitConfig() Config {
	return Config{
		MainNodeUrl: os.Getenv("MAINNODE_URL"),
		SelfUrl: os.Getenv("SELF_URL"),
		ListenPort: os.Getenv("LISTEN_PORT"),
		RootDataFolder: os.Getenv("ROOT_DATA_FOLDER"),
	}
}