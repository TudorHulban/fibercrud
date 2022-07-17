package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

type AuthorizerByIP struct{}

const (
	_urlService   = "https://ipapi.co/"
	_routeService = "/json"
)

const _authorizedISO = "CYP"

var _ Authorizer = &AuthorizerByIP{}

func (AuthorizerByIP) isValidIP(_ string) error {
	// TODO: logic

	return nil
}

// IsAuthorized returns true if IP from Cyprus using ipapi.co.
// true for  "country_code_iso3": "CYP",
func (a AuthorizerByIP) IsAuthorized(ip any) (bool, error) {
	ipRequest := ip.(string)

	if errValid := a.isValidIP(ipRequest); errValid != nil {
		return false, errValid
	}

	urlRequest := _urlService + ipRequest + _routeService

	resp, errGet := http.Get(urlRequest)
	if errGet != nil {
		return false, errGet
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return false, errGet
	}

	iso3 := gjson.Get(string(body), "country_code_iso3").String()
	if iso3 == _authorizedISO {
		return true, nil
	}

	return false, fmt.Errorf("IP: %s from country code: %s", ipRequest, iso3)
}
