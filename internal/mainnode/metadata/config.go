package metadata

import "os"

type Config struct {
	SnapShotFile string
}

func GetConfig() Config {
	return Config{
		SnapShotFile: os.Getenv("SNAPSHOT_FILE"),
	}
}