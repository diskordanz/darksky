package integration

import (
	"context"
	"time"

	"github.com/diskordanz/darksky/config"
	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (ds *Darksky) GetWeekWeather(ctx context.Context, req *api.WeekWeatherRequest) (*api.WeekWeatherResponse, error) {
	request := darksky.ForecastRequest{
		Latitude:  darksky.Measurement(req.Latitude),
		Longitude: darksky.Measurement(req.Longitude),
		Options: darksky.ForecastRequestOptions{
			Exclude: config.WeekBlocksExclude,
		},
	}

	response, err := ds.DarkskyClient.Forecast(request)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var weather []*api.DateWeatherResponse

	for _, k := range response.Daily.Data {
		weather = append(weather, &api.DateWeatherResponse{
			k.Summary,
			float64(k.TemperatureMin),
			float64(k.TemperatureMax),
			float64(k.WindSpeed),
			int(100 * k.Humidity),
			float64(k.Visibility),
			time.Unix(int64(k.Time), 0).Format(config.DateFormat),
		})
	}

	return &api.WeekWeatherResponse{
		Weather:   weather,
		WeekStart: time.Unix(int64(response.Daily.Data[0].Time), 0).Format(config.DateFormat),
		WeekEnd:   time.Unix(int64(response.Daily.Data[7].Time), 0).Format(config.DateFormat),
	}, nil
}
