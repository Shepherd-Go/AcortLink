package handler_test

import (
	"acortlink/cmd/api/handler"
	"acortlink/core/domain/models"
	mocks "acortlink/mocks/acortlink/core/domain/ports"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	ctx = context.Background()

	dataCreateUrlIsNull = models.URLCreate{
		URL:  "",
		Path: "test",
	}

	dataCreateIsGood = models.URLCreate{
		URL: "http://test.com",
	}

	dataSearchIsEmpty = models.Path{
		Path: "",
	}

	dataSearchIsGood = models.Path{
		Path: "test",
	}
)

func TestShortenerSuite(t *testing.T) {
	suite.Run(t, new(ShortenerSuiteTest))
}

type ShortenerSuiteTest struct {
	suite.Suite
	service   *mocks.ShortenApp
	underTest handler.ShortenerRequest
}

func (suite *ShortenerSuiteTest) SetupTest() {
	suite.service = mocks.NewShortenApp(suite.T())
	suite.underTest = handler.NewShortener(suite.service)
}

func (suite *ShortenerSuiteTest) TestCreate_WhenBindFail() {

	body, _ := json.Marshal("")
	controller := SetupControllerCase("/api/create", http.MethodPost, bytes.NewBuffer(body))

	err := suite.underTest.CreateShortURL(controller.Ctx)
	suite.Contains(err.Error(), "400")
	suite.Error(err)
}

func (suite *ShortenerSuiteTest) TestCreate_WhenUrlIsNull() {

	body, _ := json.Marshal(dataCreateUrlIsNull)
	controller := SetupControllerCase("/api/create", http.MethodPost, bytes.NewBuffer(body))

	err := suite.underTest.CreateShortURL(controller.Ctx)
	suite.Contains(err.Error(), "422")
	suite.Error(err)
}

func (suite *ShortenerSuiteTest) TestCreate_WhenServiceFail() {

	body, _ := json.Marshal(dataCreateIsGood)
	controller := SetupControllerCase("/api/create", http.MethodPost, bytes.NewBuffer(body))

	suite.service.Mock.On("CreateShortURL", ctx, dataCreateIsGood).
		Return("", errors.New("Error"))

	suite.Error(suite.underTest.CreateShortURL(controller.Ctx))

}

func (suite *ShortenerSuiteTest) TestCreate_WhenSuccess() {

	body, _ := json.Marshal(dataCreateIsGood)
	controller := SetupControllerCase("/api/create", http.MethodPost, bytes.NewBuffer(body))

	suite.service.Mock.On("CreateShortURL", ctx, dataCreateIsGood).
		Return("http://test.link/test", nil)

	suite.NoError(suite.underTest.CreateShortURL(controller.Ctx))

}

func (suite *ShortenerSuiteTest) TestSearch_WhenBindFail() {

	body, _ := json.Marshal("")
	controller := SetupControllerCase("/api/search?path=", http.MethodGet, bytes.NewBuffer(body))

	err := suite.underTest.SearchOriginalUrl(controller.Ctx)
	suite.Contains(err.Error(), "400")
	suite.Error(err)

}

func (suite *ShortenerSuiteTest) TestSearch_WhenPathIsEmpty() {

	body, _ := json.Marshal(dataSearchIsEmpty)
	controller := SetupControllerCase("/api/search?path=", http.MethodGet, bytes.NewBuffer(body))

	err := suite.underTest.SearchOriginalUrl(controller.Ctx)
	suite.Contains(err.Error(), "422")
	suite.Error(err)

}

func (suite *ShortenerSuiteTest) TestSearch_WhenServiceFail() {

	body, _ := json.Marshal(dataSearchIsGood)
	controller := SetupControllerCase("/api/search?path=test", http.MethodGet, bytes.NewBuffer(body))

	suite.service.Mock.On("SearchUrl", ctx, dataSearchIsGood.Path).
		Return("", errors.New("Error"))

	suite.Error(suite.underTest.SearchOriginalUrl(controller.Ctx))

}

func (suite *ShortenerSuiteTest) TestSearch_WhenSuccess() {

	body, _ := json.Marshal(dataSearchIsGood)
	controller := SetupControllerCase("/api/search?path=test", http.MethodGet, bytes.NewBuffer(body))

	suite.service.Mock.On("SearchUrl", ctx, dataSearchIsGood.Path).
		Return("http://test.com", nil)

	suite.NoError(suite.underTest.SearchOriginalUrl(controller.Ctx))

}
