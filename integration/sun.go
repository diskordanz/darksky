package integration

import (
	"context"
	"time"

	"github.com/diskordanz/darksky/config"
	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
	"github.com/xedinaska/int-weather-sdk/api"
)

func (ds *Darksky) GetSunInfo(ctx context.Context, req *api.SunInfoRequest) (*api.SunInfoResponse, error) {
	t, err := time.Parse(config.DateFormat, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	request := darksky.ForecastRequest{
		Latitude:  darksky.Measurement(req.Latitude),
		Longitude: darksky.Measurement(req.Longitude),
		Time:      darksky.Timestamp(t.Unix()),
		Options: darksky.ForecastRequestOptions{
			Exclude: config.SunBlocksExclude,
		},
	}

	response, err := ds.DarkskyClient.Forecast(request)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	return &api.SunInfoResponse{
		Rise: time.Unix(int64(response.Daily.Data[0].SunriseTime), 0),
		Set:  time.Unix(int64(response.Daily.Data[0].SunsetTime), 0),
	}, nil
}
