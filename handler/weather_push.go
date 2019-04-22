package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	sdk "github.com/xedinaska/int-weather-sdk"
	"github.com/xedinaska/int-weather-sdk/api"
)

type WeatherPushHandler struct {
	serviceName string
	logger      *logrus.Entry
	integration sdk.Integration
}

func NewWeatherPushHandler(serviceName string, logger *logrus.Entry, integration sdk.Integration) *WeatherPushHandler {
	return &WeatherPushHandler{
		serviceName: serviceName,
		logger:      logger,
		integration: integration,
	}
}

func (h *WeatherPushHandler) GetCurrentWeather(req *restful.Request, rsp *restful.Response) {
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		h.logger.Errorf("failed to unmarshal request body %v", err.Error())
		return
	}
	defer req.Request.Body.Close()

	currentWeather := &api.TodayWeatherResponse{}
	if err := json.Unmarshal(body, currentWeather); err != nil {
		h.logger.WithField("request", fmt.Sprintf("%+v", req)).Errorf("[%s] invalid incoming UpdateWeather request", h.serviceName)
		rsp.WriteHeaderAndEntity(422, &api.Error{
			Status:  422,
			Message: err.Error(),
		})
		return
	}

	fmt.Println(fmt.Sprintf("%v\n", h.formatResponse(currentWeather)))

	rsp.WriteHeaderAndEntity(http.StatusOK, currentWeather)
}

func (h *WeatherPushHandler) formatResponse(response interface{}) string {
	s, err := json.Marshal(response)
	if err != nil {
		h.logger.Fatal(err.Error())
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, s, "", "  ")

	return string(prettyJSON.Bytes())
}
