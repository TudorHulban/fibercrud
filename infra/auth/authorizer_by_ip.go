package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TudorHulban/fibercrud/infra"
	"github.com/tidwall/gjson"
)

type AuthorizerByIP struct {
	urlService string
}

const (
	_urlIPApi     = "https://ipapi.co"
	_routeService = "/json"
)

const _authorizedISO = "CYP"

var errCorruptedResponse = errors.New("response is corrupted")

var _ infra.Authorizer = &AuthorizerByIP{}

func newAuthorizerByIP(urlService string) *AuthorizerByIP {
	return &AuthorizerByIP{
		urlService: urlService,
	}
}

func NewAuthorizerByIPApi() *AuthorizerByIP {
	return &AuthorizerByIP{
		urlService: _urlIPApi,
	}
}

func (AuthorizerByIP) isValidIP(_ string) error {
	// TODO: logic
	return nil
}

// IsAuthorized returns true if IP from Cyprus using ipapi.co.
// true for  "country_code_iso3": "CYP".
func (a AuthorizerByIP) IsAuthorized(ip any) (bool, error) {
	ipRequest := ip.(string)

	if errValid := a.isValidIP(ipRequest); errValid != nil {
		return false, errValid
	}

	urlRequest := a.urlService + "/" + ipRequest + _routeService

	resp, errGet := http.Get(urlRequest)
	if errGet != nil {
		return false, errGet
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return false, errGet
	}

	if len(body) == 0 {
		return false, fmt.Errorf("IP location service: %w", errCorruptedResponse)
	}

	iso3 := gjson.Get(string(body), "country_code_iso3").String()
	if iso3 == _authorizedISO {
		return true, nil
	}

	return false, fmt.Errorf("IP: %s from country code: %s", ipRequest, iso3)
}
