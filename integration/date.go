package integration

import (
	"context"
	"time"

	"github.com/diskordanz/darksky/config"
	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (ds *Darksky) GetWeatherForDate(ctx context.Context, req *api.DateWeatherRequest) (*api.DateWeatherResponse, error) {
	t, err := time.Parse(config.DateFormat, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	request := darksky.ForecastRequest{
		Latitude:  darksky.Measurement(req.Latitude),
		Longitude: darksky.Measurement(req.Longitude),
		Time:      darksky.Timestamp(t.Unix()),
		Options: darksky.ForecastRequestOptions{
			Exclude: DateBlocksExclude,
		},
	}

	response, err := ds.DarkskyClient.Forecast(request)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	return &api.DateWeatherResponse{
		StateName:  response.Currently.Summary,
		MinTemp:    float64(response.Daily.Data[0].TemperatureMin),
		MaxTemp:    float64(response.Daily.Data[0].TemperatureMax),
		WindSpeed:  float64(response.Currently.WindSpeed),
		Humidity:   int(100 * response.Currently.Humidity),
		Visibility: float64(response.Currently.Visibility),
		Date:       time.Unix(int64(response.Currently.Time), 0).Format(config.DateFormat),
	}, nil
}
