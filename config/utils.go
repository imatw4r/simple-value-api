package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var globalConfig AppConfig
var configOnce sync.Once

func GetConfig() AppConfig {
	configOnce.Do(func() {
		globalConfig = loadConfigFromEnv()
		InitLog(globalConfig)
	})
	return globalConfig
}

func loadConfigFromEnv() AppConfig {
	godotenv.Load()
	port := os.Getenv("PORT")
	percentageValue := os.Getenv("VALUE_ACCEPTABLE_DIFF_PERCENTAGE")
	logLevel := os.Getenv("LOG_LEVEL")
	sourceFilePath := os.Getenv("SOURCE_FILEPATH")
	acceptableValueDiffPercentage, err := strconv.ParseFloat(percentageValue, 64)

	if err != nil {
		panic("Failed to load values from environment.")
	}

	return NewConfig(
		port, acceptableValueDiffPercentage, logLevel, sourceFilePath,
	)
}
