package integration

import (
	"context"

	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (i *Example) GetSunInfo(ctx context.Context, req *api.SunInfoRequest) (*api.SunInfoResponse, error) {
	return nil, errors.New("not implemented")
}
