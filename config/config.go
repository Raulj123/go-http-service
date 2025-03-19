package config

import (
	"encoding/json"
	"os"
)
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`

	DBname string `json:"dbname"`
	Driver string `json:"driver"`
	DBuri string `json:"DBuri"`
	// for now this is very basic
}

// This function reads a env file in your project root and loads the configuration value
func LoadConfig(path string) (*Config, error){
	myConfig := Config{}
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configFile, &myConfig)
	if err != nil {
		return nil, err
	}
	return &myConfig, nil
}


