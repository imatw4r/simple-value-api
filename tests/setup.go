package tests

import (
	"value-app/config"
)

func GetTestConfig() config.AppConfig {
	return config.AppConfig{
		Port:                          "8080",
		LogLevel:                      "DEBUG",
		AcceptableValueDiffPercentage: 10,
		SourceFilepath:                "test.in",
	}
}
