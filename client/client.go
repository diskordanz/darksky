package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/diskordanz/darksky/config"
	"github.com/diskordanz/darksky/integration"
	log "github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk/api"
)

var ctx context.Context
var logger *log.Entry
var example *integration.Example

func main() {
	serviceName, serviceVersion := "darksky", "0.0.1"

	ctx = context.Background()
	ctx = context.WithValue(ctx, config.BaseURL, os.Getenv("API_URL"))

	logger = log.WithFields(log.Fields{
		"logger": "main",
		"serviceContext": map[string]string{
			"service": serviceName,
			"version": serviceVersion,
		},
	})

	example = integration.Init(logger)

	getWeatherToday()
	getWeekWeather()
	//todo: add other calls
}

func getWeatherToday() {
	response, err := example.GetTodayWeather(ctx, &api.TodayWeatherRequest{
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%v\n", formatResponse(response)))
}

func getWeekWeather() {
	response, err := example.GetWeekWeather(ctx, &api.WeekWeatherRequest{
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%v\n", formatResponse(response)))
}

func formatResponse(response interface{}) string {
	s, err := json.Marshal(response)
	if err != nil {
		logger.Fatal(err.Error())
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, s, "", "  ")

	return string(prettyJSON.Bytes())
}
