package integration

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/diskordanz/darksky/config"
	"github.com/diskordanz/darksky/request"
	"github.com/shawntoffel/darksky"
	"github.com/sirupsen/logrus"
)

const (
	TodayBlocksExclude = "hourly, minutely, alerts, flags"
	DateBlocksExclude  = "hourly, minutely, alerts, flags"
	WeekBlocksExclude  = "currently, hourly, minutely, alerts, flags"
	SunBlocksExclude   = "currently, hourly, minutely, alerts, flags"
)

type Darksky struct {
	Logger        *logrus.Entry
	RequestClient *request.RequestClient
	DarkskyClient darksky.DarkSky
}

func Init(logger *logrus.Entry) *Darksky {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	key := os.Getenv(config.APIKey)
	if key == "" {
		logger.Fatalf("%s env variable should be provided", config.APIKey)
	}

	darkskyClient := darksky.NewWithClient(key, httpClient)

	return &Darksky{
		Logger: logger,
		RequestClient: &request.RequestClient{
			HTTPClient: httpClient,
		},
		DarkskyClient: darkskyClient,
	}
}
