package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/diskordanz/darksky/integration"
	log "github.com/sirupsen/logrus"
	"github.com/xedinaska/int-weather-sdk/api"
)

var ctx context.Context
var logger *log.Entry
var darksky *integration.Darksky

func main() {
	serviceName, serviceVersion := "darksky", "0.0.1"

	ctx = context.Background()

	logger = log.WithFields(log.Fields{
		"logger": "main",
		"serviceContext": map[string]string{
			"service": serviceName,
			"version": serviceVersion,
		},
	})

	darksky = integration.Init(logger)

	getWeatherToday()
	getWeekWeather()
	getWeatherForDate()
	getSunInfo()

}

func getWeatherToday() {
	response, err := darksky.GetTodayWeather(ctx, &api.TodayWeatherRequest{
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%v\n", formatResponse(response)))
}

func getWeekWeather() {
	response, err := darksky.GetWeekWeather(ctx, &api.WeekWeatherRequest{
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%v\n", formatResponse(response)))
}

func getSunInfo() {
	response, err := darksky.GetSunInfo(ctx, &api.SunInfoRequest{
		Date:      "2019-04-20",
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%v\n", formatResponse(response)))
}

func getWeatherForDate() {
	response, err := darksky.GetWeatherForDate(ctx, &api.DateWeatherRequest{
		Date:      "2019-04-20",
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
