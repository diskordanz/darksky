package integration

import (
	"crypto/tls"
	"net/http"

	"github.com/diskordanz/darksky/request"
	"github.com/shawntoffel/darksky"
	"github.com/sirupsen/logrus"
)

type Example struct {
	Logger        *logrus.Entry
	RequestClient *request.RequestClient
	DarkskyClient darksky.DarkSky
}

func Init(logger *logrus.Entry) *Example {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	darkskyClient := darksky.NewWithClient("b04ad8db6f75cbd1a02e6e4c8e1e1272", httpClient)

	return &Example{
		Logger: logger,
		RequestClient: &request.RequestClient{
			HTTPClient: httpClient,
		},
		DarkskyClient: darkskyClient,
	}
}
