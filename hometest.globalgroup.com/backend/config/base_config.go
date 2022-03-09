package config

import (
	"os"
)

func GetConfigFileName() string {
	dir, _ := os.Getwd()
	return dir + "/hometest.globalgroup.com/backend/config/base_config.json"
}
