package integration_test

import (
	"testing"

	"github.com/shawntoffel/darksky"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type DateWeatherTestSuite struct {
	IntegrationTestSuite
}

func (s *DateWeatherTestSuite) TestGetDateWeather_Failure() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/date_weather_failed.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	_, err = s.app.GetWeatherForDate(s.ctx, &api.DateWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
		Date:      "2019-04-20",
	})

	s.Require().NotNil(err)
	s.Require().Equal("response from server isn't success", err.Error())
	s.TeardownTest()
}

func (s *DateWeatherTestSuite) TestGetDateWeather_Success() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/date_weather_success.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	response, err := s.app.GetWeatherForDate(s.ctx, &api.DateWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
		Date:      "2019-04-20",
	})

	s.Require().NoError(err)
	s.Require().Equal("2019-01-19", response.Date)
	s.Require().Equal("Clear", response.StateName)
	s.Require().Equal(28.78, response.MinTemp)
	s.Require().Equal(37.65, response.MaxTemp)
	s.Require().Equal(70, response.Humidity)
	s.Require().Equal(4.46, response.WindSpeed)
	s.Require().Equal(10.0, response.Visibility)
	s.TeardownTest()
}

func TestDateWeather(t *testing.T) {
	suite.Run(t, new(DateWeatherTestSuite))
}
