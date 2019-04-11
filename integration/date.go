package integration

import (
	"context"

	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (i *Example) GetWeatherForDate(ctx context.Context, req *api.DateWeatherRequest) (*api.DateWeatherResponse, error) {
	return nil, errors.New("not implemented")
}
