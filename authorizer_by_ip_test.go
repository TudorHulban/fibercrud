package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_ipNicosia = "45.89.23.105"
	_ipRomania = "82.77.245.236"
)

func TestAuthorizerByIP(t *testing.T) {
	var auth AuthorizerByIP

	isAuthorized, errAuth := auth.IsAuthorized(_ipNicosia)
	require.NoError(t, errAuth)
	require.True(t, isAuthorized)

	isNotAuthorized, errNoAuth := auth.IsAuthorized(_ipRomania)
	require.Error(t, errNoAuth)
	require.False(t, isNotAuthorized)
}
