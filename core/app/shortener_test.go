package app_test

import (
	"acortlink/core/app"
	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"
	mocks "acortlink/mocks/acortlink/core/domain/ports"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

var (
	ctx = context.Background()

	urlIsGood = models.URLCreate{
		URL:  "http://test.com",
		Path: "test",
	}
)

func TestShortenerSuite(t *testing.T) {
	suite.Run(t, new(ShortenerTestApp))
}

type ShortenerTestApp struct {
	suite.Suite
	postgr    *mocks.ShortenRepoPostgres
	redis     *mocks.ShortenRepoRedis
	underTest ports.ShortenApp
}

func (suite *ShortenerTestApp) SetupTest() {

	suite.postgr = &mocks.ShortenRepoPostgres{}
	suite.redis = &mocks.ShortenRepoRedis{}
	suite.underTest = app.NewShortenApp(suite.postgr, suite.redis)

}

func (suite *ShortenerTestApp) TestCreate_WhenRepoPostgresFail() {

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{}, errors.New("Error"))

	_, err := suite.underTest.CreateShortURL(ctx, urlIsGood)

	suite.Contains(err.Error(), "unexpected error")
	suite.Contains(err.Error(), "500")
	suite.Error(err)

}

func (suite *ShortenerTestApp) TestCreate_WhenPathExists() {

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{URL: "test"}, nil)

	_, err := suite.underTest.CreateShortURL(ctx, urlIsGood)

	suite.Contains(err.Error(), "path already exists")
	suite.Contains(err.Error(), "409")
	suite.Error(err)

}

func (suite *ShortenerTestApp) TestCreate_WhenCreateShortenFail() {

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{}, nil)

	suite.postgr.Mock.On("Save", ctx, urlIsGood).
		Return(errors.New("Error"))

	_, err := suite.underTest.CreateShortURL(ctx, urlIsGood)

	suite.Contains(err.Error(), "unexpected error")
	suite.Contains(err.Error(), "500")
	suite.Error(err)

}

func (suite *ShortenerTestApp) TestCreate_WhenSuccess() {

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{}, nil)

	suite.postgr.Mock.On("Save", ctx, urlIsGood).
		Return(nil)

	suite.postgr.Mock.On("AddContToQuerysUrl", ctx, uuid.New()).Return(nil)

	_, err := suite.underTest.CreateShortURL(ctx, urlIsGood)

	suite.NoError(err)

}

func (suite *ShortenerTestApp) TestSearch_WhenRepoPostgresFail() {

	suite.redis.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return("", nil)

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{}, errors.New("Error"))

	suite.postgr.Mock.On("AddContToQuerysUrl", ctx, uuid.New()).Return(nil)

	_, err := suite.underTest.SearchUrl(ctx, urlIsGood.Path)

	suite.Contains(err.Error(), "unexpected error")
	suite.Contains(err.Error(), "500")
	suite.Error(err)

}

func (suite *ShortenerTestApp) TestSearch_WhenUrlNotFound() {

	suite.redis.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return("", nil)

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{}, nil)

	_, err := suite.underTest.SearchUrl(ctx, urlIsGood.Path)

	suite.Contains(err.Error(), "url not found")
	suite.Contains(err.Error(), "404")
	suite.Error(err)

}

func (suite *ShortenerTestApp) TestSearch_WhenSuccess() {

	suite.redis.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return("", nil)

	suite.postgr.Mock.On("SearchUrl", ctx, urlIsGood.Path).
		Return(models.URLResponse{URL: "test"}, nil)

	suite.redis.Mock.On("Save", ctx, "test", "test", 0*time.Second).
		Return(nil)

	_, err := suite.underTest.SearchUrl(ctx, urlIsGood.Path)

	suite.NoError(err)

}
