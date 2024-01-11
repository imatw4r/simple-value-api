package tests

import (
	"value-app/config"
	"value-app/pkg/domain"
)

func getTestConfig() *config.GlobalConfig {
	return config.InitConfig()
}

func GetValueService(source domain.IValueSource) *domain.ValueService {
	config := getTestConfig()
	values, _ := source.Load()
	svc := domain.NewValueService(config.App.AcceptableValueDiffPercentage)
	svc.AddValues(values)
	return svc
}
