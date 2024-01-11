package test_domain

import (
	"value-app/domain"
	test "value-app/tests"
	stub "value-app/tests/stubs"
)

func getValueService() *domain.ValueService {
	config := test.GetTestConfig()

	valueSource := stub.NewValueSource()
	values, _ := valueSource.Load()
	svc := domain.NewValueService(config.App.AcceptableValueDiffPercentage)
	svc.AddValues(values)
	return svc
}
