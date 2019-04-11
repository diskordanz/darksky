package main

import (
	"github.com/diskordanz/darksky/integration"
	log "github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk/config"
	"github.com/xedinaska/int-weather-sdk/server"
)

const (
	serviceName    = "darksky"
	serviceVersion = "0.0.1"
)

var (
	logger = log.WithFields(log.Fields{
		"logger": "main",
		"serviceContext": map[string]string{
			"service": serviceName,
			"version": serviceVersion,
		},
	})
)

func main() {
	conf, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	/*apiURL := os.Getenv(intConf.BaseURL)
	if apiURL == "" {
		logger.Fatalf("%s env variable should be provided", intConf.BaseURL)
	}*/

	exampleIntegration := integration.Init(logger)
	srv := server.Create(serviceName, serviceVersion, exampleIntegration, conf)

	logger.Info("Starting service")
	if err := srv.WebService.Run(); err != nil {
		logger.Fatal(err)
	}
}
