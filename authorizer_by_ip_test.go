package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_ipNicosia       = "45.89.23.105"
	_responseNicosia = `{
		"ip": "45.89.23.105",
		"version": "IPv4",
		"city": "Nicosia",
		"region": "Nicosia",
		"region_code": "01",
		"country": "CY",
		"country_name": "Cyprus",
		"country_code": "CY",
		"country_code_iso3": "CYP",
		"country_capital": "Nicosia",
		"country_tld": ".cy",
		"continent_code": "EU",
		"in_eu": true,
		"postal": null,
		"latitude": 35.1638,
		"longitude": 33.3639,
		"timezone": "Asia/Nicosia",
		"utc_offset": "+0300",
		"country_calling_code": "+357",
		"currency": "EUR",
		"currency_name": "Euro",
		"languages": "el-CY,tr-CY,en",
		"country_area": 9250.0,
		"country_population": 1189265,
		"asn": "AS211239",
		"org": "C.V.M. WiNet Solutions LTD"
	}`

	_ipRomania       = "82.77.245.232"
	_responseRomania = `{
		"ip": "82.77.245.232",
		"version": "IPv4",
		"city": "Iasi",
		"region": "Iasi",
		"region_code": "IS",
		"country": "RO",
		"country_name": "Romania",
		"country_code": "RO",
		"country_code_iso3": "ROU",
		"country_capital": "Bucharest",
		"country_tld": ".ro",
		"continent_code": "EU",
		"in_eu": true,
		"postal": "700023",
		"latitude": 47.1672,
		"longitude": 27.6083,
		"timezone": "Europe/Bucharest",
		"utc_offset": "+0300",
		"country_calling_code": "+40",
		"currency": "RON",
		"currency_name": "Leu",
		"languages": "ro,hu,rom",
		"country_area": 237500.0,
		"country_population": 19473936,
		"asn": "AS8708",
		"org": "RCS & RDS"
	}`

	_ipCloudflare      = "1.1.1.1"
	_responseCorrupted = ""
)

func TestAuthorizerByIP(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")[0] {
		case _ipNicosia:
			fmt.Fprint(w, _responseNicosia)

		case _ipRomania:
			fmt.Fprint(w, _responseRomania)

		default:
			fmt.Fprint(w, _responseCorrupted)
		}
	}))

	defer mockServer.Close()

	auth := newAuthorizerByIP(mockServer.URL)

	isAuthorized, errAuth := auth.IsAuthorized(_ipNicosia)
	require.NoError(t, errAuth)
	require.True(t, isAuthorized)

	isNotAuthorized, errNoAuth := auth.IsAuthorized(_ipRomania)
	require.Error(t, errNoAuth)
	require.False(t, isNotAuthorized)

	isCorrupted, errCo := auth.IsAuthorized(_ipCloudflare)
	require.ErrorIs(t, errors.Unwrap(errCo), errCorruptedResponse)
	require.False(t, isCorrupted)
}
