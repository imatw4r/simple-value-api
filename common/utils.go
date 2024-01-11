package common

import (
	"os"
	"sync"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var globalConfig GlobalConfig
var configOnce sync.Once

func InitConfig() *GlobalConfig {
	configOnce.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		// Use dev config by default
		if configPath == "" {
			configPath = "./config/dev.yaml"
		}
		config := loadConfigFromFile(configPath)
		InitLog(config)
	})
	return &globalConfig
}

func loadConfigFromFile(filePath string) *GlobalConfig {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the configuration file: %v", err)
	}
	err = yaml.Unmarshal(configFile, &globalConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}
	return &globalConfig
}

func InitLog(config *GlobalConfig) {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"timestamp", "severity", "component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})
	log.SetLevel(getLoggerLevel(config.Logger.LogLevel))
}

func getLoggerLevel(value string) log.Level {
	log.Infof("Setting log level: %v", value)
	logLevel, err := log.ParseLevel(value)
	if err != nil {
		return log.InfoLevel
	}
	return logLevel
}
