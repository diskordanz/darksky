package integration

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (ds *Darksky) GetTodayWeather(ctx context.Context, req *api.TodayWeatherRequest) (*api.TodayWeatherResponse, error) {
	request := darksky.ForecastRequest{
		Latitude:  darksky.Measurement(req.Latitude),
		Longitude: darksky.Measurement(req.Longitude),
		Time:      darksky.Timestamp(time.Now().Unix()),
		Options: darksky.ForecastRequestOptions{
			Exclude: TodayBlocksExclude,
		},
	}

	response, err := ds.DarkskyClient.Forecast(request)
	if err != nil || response.Currently == nil || response.Daily == nil {
		return nil, errors.New("response from server isn't success")
	}

	return &api.TodayWeatherResponse{
		StateName:  response.Currently.Summary,
		MinTemp:    float64(response.Daily.Data[0].TemperatureMin),
		MaxTemp:    float64(response.Daily.Data[0].TemperatureMax),
		WindSpeed:  float64(response.Currently.WindSpeed),
		Humidity:   int(100 * response.Currently.Humidity),
		Visibility: float64(response.Currently.Visibility),
	}, nil
}
