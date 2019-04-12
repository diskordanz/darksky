package integration

import (
	"crypto/tls"
	"net/http"

	"github.com/diskordanz/darksky/config"
	"github.com/diskordanz/darksky/request"
	"github.com/shawntoffel/darksky"
	"github.com/sirupsen/logrus"
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

	darkskyClient := darksky.NewWithClient(config.APIKey, httpClient)

	return &Darksky{
		Logger: logger,
		RequestClient: &request.RequestClient{
			HTTPClient: httpClient,
		},
		DarkskyClient: darkskyClient,
	}
}
