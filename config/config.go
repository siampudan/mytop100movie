package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadConfig() {
	_, filePath, _, _ := runtime.Caller(0)
	configName := "config.yaml"
	configPath := filePath[:len(filePath)-9] + string(filepath.Separator)

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	setGinMode(viper.GetString("server.mode"))
}

func setGinMode(mode string) {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)
		break
	default:
		gin.SetMode(gin.DebugMode)
	}
}
