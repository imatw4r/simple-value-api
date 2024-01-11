package tests

import (
	"value-app/common"
)

func GetTestConfig() *common.GlobalConfig {
	return common.InitConfig()
}
