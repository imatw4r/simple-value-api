package main

import (
	"context"
	"value-app/api"
	"value-app/config"
	"value-app/domain"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appConfig := config.GetConfig()

	svc := domain.NewValueService(appConfig.AcceptableValueDiffPercentage)

	log.Infof("Reading source file: %s", appConfig.SourceFilepath)
	source := domain.NewFileSource(appConfig.SourceFilepath)

	values, err := source.Load()
	if err != nil {
		log.Fatalf("Failed to load file. Error %v", err)
	}
	err = svc.AddValues(values)
	if err != nil {
		log.Error("Failed to load values to the service.")
		return
	}

	log.Infof("Initiating web application")
	api.RunWebserver(ctx, cancel, appConfig, svc)
	<-ctx.Done()
	log.Info("Exiting...")
}
