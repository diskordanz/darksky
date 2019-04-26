package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/diskordanz/darksky/config"
	"github.com/diskordanz/darksky/handler"
	"github.com/diskordanz/darksky/integration"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type HandlerTestSuite struct {
	suite.Suite
	app *integration.Darksky
}

func (s *HandlerTestSuite) SetupTest() {
	os.Setenv(config.APIKey, "testApi")
	logger := logrus.WithFields(logrus.Fields{
		"logger": "handler_test",
	})
	s.app = integration.Init(logger)
}

func (s *HandlerTestSuite) TestReceiveCurrentWeather_Success() {
	s.SetupTest()
	weatherPushHandler := handler.NewWeatherPushHandler("test", s.app.Logger, s.app)

	data, err := ioutil.ReadFile("testdata/current_wether_success.json")
	s.Require().NoError(err)

	httpWriter := httptest.NewRecorder()
	httpRequest, err := http.NewRequest("POST", "testUrl", bytes.NewReader(data))
	s.Require().NoError(err)

	request := restful.NewRequest(httpRequest)
	response := restful.NewResponse(httpWriter)
	response.SetRequestAccepts(restful.MIME_JSON)

	weatherPushHandler.ReceiveCurrentWeather(request, response)

	receivedWeather := &api.TodayWeatherResponse{}
	err = json.NewDecoder(httpWriter.Body).Decode(receivedWeather)

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, response.StatusCode())
	s.Require().Equal("Clear", receivedWeather.StateName)
	s.Require().Equal(28.78, receivedWeather.MinTemp)
	s.Require().Equal(37.65, receivedWeather.MaxTemp)
	s.Require().Equal(70, receivedWeather.Humidity)
	s.Require().Equal(4.46, receivedWeather.WindSpeed)
	s.Require().Equal(10.0, receivedWeather.Visibility)
}

func (s *HandlerTestSuite) TestReceiveCurrentWeather_Failure() {
	s.SetupTest()
	weatherPushHandler := handler.NewWeatherPushHandler("test", s.app.Logger, s.app)

	data, err := ioutil.ReadFile("testdata/current_wether_failed.json")
	s.Require().NoError(err)

	httpWriter := httptest.NewRecorder()
	httpRequest, err := http.NewRequest("POST", "testUrl", bytes.NewReader(data))
	s.Require().NoError(err)

	request := restful.NewRequest(httpRequest)
	response := restful.NewResponse(httpWriter)
	response.SetRequestAccepts(restful.MIME_JSON)

	weatherPushHandler.ReceiveCurrentWeather(request, response)

	s.Require().Equal(http.StatusUnprocessableEntity, response.StatusCode())
}

func TestWeatherPushHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
