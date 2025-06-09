package fsManager

import (
	"encoding/json"
	"os"
)

type Config struct {
	Paths []string `json:"paths"`
}

func InitConfig() ([]string, error) {
	_, err := os.Stat("./config.json")
	if os.IsNotExist(err) {
		config := Config{
			Paths: []string{},
		}
		data, err := json.Marshal(config)
		if err != nil {
			return nil, err
		}
		os.WriteFile("config.json", data, 0644)
	}
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	json.Unmarshal(data, config)
	return config.Paths, nil
}
