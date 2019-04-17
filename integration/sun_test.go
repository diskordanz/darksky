package integration_test

import (
	"testing"
	"time"

	"github.com/shawntoffel/darksky"
	"github.com/stretchr/testify/suite"
	"github.com/xedinaska/int-weather-sdk/api"
)

type SunInfoTestSuite struct {
	IntegrationTestSuite
}

func (s *SunInfoTestSuite) TestGetSunInfo_Failure() {
	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/sun_info_failed.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	_, err = s.app.GetSunInfo(s.ctx, &api.SunInfoRequest{
		Latitude:  42.21,
		Longitude: 25.11,
		Date:      "2019-04-20",
	})

	s.Require().NotNil(err)
	s.Require().Equal("response from server isn't success", err.Error())
	s.TeardownTest()

}

func (s *SunInfoTestSuite) TestGetSunInfo_Success() {

	s.SetupTest()
	routeMap := map[string]string{
		"/": "testdata/sun_info_success.json",
	}
	err := s.configureAndStartTestServer(routeMap)
	s.Require().NoError(err)

	darksky.BaseUrl = s.ts.URL
	response, err := s.app.GetSunInfo(s.ctx, &api.SunInfoRequest{
		Latitude:  42.21,
		Longitude: 25.11,
		Date:      "2019-04-20",
	})

	s.Require().NoError(err)
	s.Require().Equal(time.Unix(int64(1547900227), 0), response.Rise)
	s.Require().Equal(time.Unix(int64(1547935114), 0), response.Set)

	s.TeardownTest()
}

func TestSunInfo(t *testing.T) {
	suite.Run(t, new(SunInfoTestSuite))
}
