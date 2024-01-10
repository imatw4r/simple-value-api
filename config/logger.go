package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLog(config AppConfig) {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"timestamp", "severity", "component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})
	log.SetLevel(getLoggerLevel(config.LogLevel))
}

func getLoggerLevel(value string) log.Level {
	log.Infof("Setting log level: %v", value)
	logLevel, err := log.ParseLevel(value)
	if err != nil {
		return log.InfoLevel
	}
	return logLevel
}
