package utils

import "os"

func LoadConfigFile() string {
	var configType, configFile string
	configType = os.Getenv("config")
	if configType == "docker" {
		configFile = "./config/config-docker"
	} else {
		configFile = "./config/config-local"
	}

	return configFile
}
