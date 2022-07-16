package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/require"
)

const dataCreate = "{\"code\": \"J1234\", \"name\": \"avata\", \"country\": \"Fidji\", \"website\": \"avata.fj\", \"phone\": \"+55 12345\"}"

func TestFiber(t *testing.T) {
	initialization()

	require := require.New(t)

	repo, errNew := NewRepoCompany()
	require.NoError(errNew)

	repo.Migration(&CompanyData{})

	fiber := NewFiber(_portFiber, repo)
	defer fiber.Stop()

	resp, err := fiber.app.Test(httptest.NewRequest(http.MethodGet, _route+"/1", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 404, resp.StatusCode)

	resp, err = fiber.app.Test(httptest.NewRequest(http.MethodPost, _route, strings.NewReader(dataCreate)))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)
}
