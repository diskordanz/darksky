package main

import (
	"os"
	"strconv"

	"github.com/diskordanz/darksky/handler"
	"github.com/diskordanz/darksky/integration"
	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"
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

	darksky := integration.Init(logger)

	srv := server.Create(serviceName, serviceVersion, darksky, conf)

	allowPush, err := strconv.ParseBool(os.Getenv("ENABLE_WEATHER_PUSH_REQUESTS"))
	if err != nil {
		logger.Fatal(errors.Wrap(err, "unable to parse ENABLE_WEATHER_PUSH_REQUESTS env var"))
	}

	if allowPush {
		weatherPushHandler := handler.NewWeatherPushHandler(serviceName, logger, darksky)
		if err != nil {
			logger.Fatal(errors.Wrap(err, "WeatherPushHandler initialization failed"))
		}

		restfulService := new(restful.WebService).
			Path("/").
			Consumes(restful.MIME_JSON).
			Produces(restful.MIME_JSON)

		restfulService.Route(restfulService.POST("/weather/update").To(weatherPushHandler.ReceiveCurrentWeather))
		srv.WebRouter.Add(restfulService)
	}

	logger.Info("Starting service")
	if err := srv.WebService.Run(); err != nil {
		logger.Fatal(err)
	}
}
