package main

import (
	"context"
	"value-app/config"
	api "value-app/pkg/api"
	"value-app/pkg/domain"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	setupTerminationAtSignals(ctx, cancel)

	appConfig := config.InitConfig()
	svc := initService(appConfig)

	api.RunWebserver(ctx, cancel, appConfig, svc)
	<-ctx.Done()
}

func initService(config *config.GlobalConfig) *domain.ValueService {
	svc := domain.NewValueService(config.App.AcceptableValueDiffPercentage)
	log.Infof("Reading source file: %s", config.App.SourceFilepath)
	source := domain.NewFileSource(config.App.SourceFilepath)

	values, err := source.Load()
	if err != nil {
		log.Fatalf("Failed to load file. Error %v", err)
	}
	err = svc.AddValues(values)
	if err != nil {
		log.Fatalf("Failed to load values to the service.")
	}
	return svc
}
