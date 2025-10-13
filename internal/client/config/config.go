package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	MetadataServer     string `json:"metadata_server"`
	DefaultConcurrency int    `json:"default_concurrency"`
}

func defaultConfig() Config {
	return Config{
		MetadataServer: "http://localhost:8080",
		DefaultConcurrency: 4,
	}
}

func LoadConfig() (Config, error) {
	path, err := getConfigPath()

	if err != nil {
		return Config {}, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfg := defaultConfig()
		data, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return Config{}, err
		}
		if err := os.WriteFile(path, data, os.FileMode(os.O_CREATE)); err != nil {
			return Config{}, fmt.Errorf("Failed to create deafult config with error %w", err)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file : %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}

func getConfigPath() (string, error) {
	home_dir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Failed to get Home dir with error %w", err)
	}
	dir := filepath.Join(home_dir, "dfsctl")
	dir_err := os.MkdirAll(dir, os.ModeDir)

	if dir_err != nil {
		return  "", fmt.Errorf("Failed to create dir with error %w", err)
	}
	
	return filepath.Join(dir, "config.json"), nil
}