package integration

import (
	"context"
	"log"

	"github.com/shawntoffel/darksky"
	"github.com/xedinaska/int-weather-sdk/api"
)

func UpdateWeather(ctx context.Context, req *darksky.ForecastResponse) (*api.TodayWeatherResponse, error) {

	log.Println(req.Currently.Summary)

	return &api.TodayWeatherResponse{
		StateName:  req.Currently.Summary,
		MinTemp:    float64(req.Daily.Data[0].TemperatureMin),
		MaxTemp:    float64(req.Daily.Data[0].TemperatureMax),
		WindSpeed:  float64(req.Currently.WindSpeed),
		Humidity:   int(100 * req.Currently.Humidity),
		Visibility: float64(req.Currently.Visibility),
	}, nil
}
