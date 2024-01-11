package tests

import (
	"value-app/common"
	"value-app/domain"
	stub "value-app/tests/stubs"
)

func getTestConfig() *common.GlobalConfig {
	return common.InitConfig()
}

func GetValueService() *domain.ValueService {
	config := getTestConfig()

	valueSource := stub.NewValueSource()
	values, _ := valueSource.Load()
	svc := domain.NewValueService(config.App.AcceptableValueDiffPercentage)
	svc.AddValues(values)
	return svc
}
