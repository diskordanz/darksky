package integration_test

import (
	"testing"

	"github.com/shawntoffel/darksky"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type TodayWeatherTestSuite struct {
	IntegrationTestSuite
}

func (s *TodayWeatherTestSuite) TestGetTodayWeather_Failure() {
	routeMap := map[string]string{
		"/": "testdata/today_weather_failed.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	_, err = s.app.GetTodayWeather(s.ctx, &api.TodayWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
	})

	s.Require().NotNil(err)
	s.Require().Equal("response from server isn't success", err.Error())
}

func (s *TodayWeatherTestSuite) TestGetTodayWeather_Success() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/today_weather_success.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	response, err := s.app.GetTodayWeather(s.ctx, &api.TodayWeatherRequest{
		Latitude:  53.9045,
		Longitude: 27.5615,
	})

	s.Require().NoError(err)
	s.Require().Equal("Clear", response.StateName)
	s.Require().Equal(28.78, response.MinTemp)
	s.Require().Equal(37.65, response.MaxTemp)
	s.Require().Equal(70, response.Humidity)
	s.Require().Equal(4.46, response.WindSpeed)
	s.Require().Equal(10.0, response.Visibility)
	s.TeardownTest()
}

func TestTodayWeather(t *testing.T) {
	suite.Run(t, new(TodayWeatherTestSuite))
}
