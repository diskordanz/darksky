package integration

import (
	"context"

	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (i *Example) GetWeekWeather(ctx context.Context, req *api.WeekWeatherRequest) (*api.WeekWeatherResponse, error) {
	return nil, errors.New("not implemented")
}
