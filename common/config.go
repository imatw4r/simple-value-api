package common

type LoggerConfig struct {
	LogLevel string `yaml:"logLevel"`
}

type AppConfig struct {
	Port                          string  `yaml:"port"`
	AcceptableValueDiffPercentage float64 `yaml:"acceptableValueDiffPercentage"`
	SourceFilepath                string  `yaml:"sourceFilepath"`
}

type GlobalConfig struct {
	Logger LoggerConfig `yaml:"logging"`
	App    AppConfig    `yaml:"app"`
}
