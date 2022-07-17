package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TudorHulban/fibercrud/domain"
	"github.com/TudorHulban/fibercrud/infra"
	repo "github.com/TudorHulban/fibercrud/repository"
	"github.com/tidwall/gjson"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/require"
)

const dataCreate = `{"code":"J1234","name":"avata","country":"Fidji","website":"avata.fj","phone":"+55 12345"}`

func TestFiber(t *testing.T) {
	infra.Initialization()

	require := require.New(t)

	repo, errNew := repo.NewRepoCompany()
	require.NoError(errNew)

	repo.Migration(&domain.CompanyData{})

	fiber := NewFiber(repo)
	defer fiber.Stop()

	fiber.addRoutes()

	resp, err := fiber.app.Test(httptest.NewRequest(http.MethodGet, _route+"/1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 404, resp.StatusCode)

	req := httptest.NewRequest(http.MethodPost, _route, strings.NewReader(dataCreate))
	req.Header.Set("Content-type", "application/json")

	resp, err = fiber.app.Test(req)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)

	resp, err = fiber.app.Test(httptest.NewRequest(http.MethodGet, _route+"/1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	require.NoError(errRead)
	require.Equal("J1234", gjson.Get(string(body), "company.code").String())
	require.Equal("avata", gjson.Get(string(body), "company.name").String())
	require.Equal("Fidji", gjson.Get(string(body), "company.country").String())
	require.Equal("avata.fj", gjson.Get(string(body), "company.website").String())
	require.Equal("+55 12345", gjson.Get(string(body), "company.phone").String())

	resp, err = fiber.app.Test(httptest.NewRequest(http.MethodDelete, _route+"/1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)

	resp, err = fiber.app.Test(httptest.NewRequest(http.MethodGet, _route+"/1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 404, resp.StatusCode)

	resp, err = fiber.app.Test(httptest.NewRequest(http.MethodGet, _route+"/-1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 400, resp.StatusCode)
}
