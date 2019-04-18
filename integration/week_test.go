package integration_test

import (
	"testing"

	"github.com/shawntoffel/darksky"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type WeekWeatherTestSuite struct {
	IntegrationTestSuite
}

func (s *WeekWeatherTestSuite) TestGetWeekWeather_Failure() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/week_weather_failed.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	_, err = s.app.GetWeekWeather(s.ctx, &api.WeekWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
	})

	s.Require().NotNil(err)
	s.Require().Equal("response from server isn't success", err.Error())
	s.TeardownTest()
}

func (s *WeekWeatherTestSuite) TestGetWeekWeather_Success() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/week_weather_success.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	response, err := s.app.GetWeekWeather(s.ctx, &api.WeekWeatherRequest{
		Latitude:  42.21,
		Longitude: 25.11,
	})

	s.Require().NoError(err)
	s.Require().Equal("2019-01-19", response.WeekStart)
	s.Require().Equal("2019-01-26", response.WeekEnd)

	s.Require().Equal("Clear", response.Weather[0].StateName)
	s.Require().Equal(28.78, response.Weather[0].MinTemp)
	s.Require().Equal(37.65, response.Weather[0].MaxTemp)
	s.Require().Equal(74, response.Weather[0].Humidity)
	s.Require().Equal(4.4, response.Weather[0].WindSpeed)
	s.Require().Equal(10.0, response.Weather[0].Visibility)
	s.TeardownTest()
}

func TestWeekWeather(t *testing.T) {
	suite.Run(t, new(WeekWeatherTestSuite))
}
