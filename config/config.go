package config

type AppConfig struct {
	Port                          string
	AcceptableValueDiffPercentage float64
	LogLevel                      string
	SourceFilepath                string
}

func NewConfig(
	port string,
	acceptableValueDiffPercentage float64,
	logLevel string,
	sourceFilepath string,
) AppConfig {
	return AppConfig{
		Port:                          port,
		AcceptableValueDiffPercentage: acceptableValueDiffPercentage,
		LogLevel:                      logLevel,
		SourceFilepath:                sourceFilepath,
	}
}
