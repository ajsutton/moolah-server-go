package main

import (
	"github.com/moolah-server-go/infrastructure/services/web"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListAccounts(t *testing.T) {
	application := NullApplication()
	application.RegisterHandlers()

	status, got, err := application.router.Call(web.CallData{
		Method: http.MethodGet,
		Url:    "/api/accounts/",
	})
	require.Nil(t, err, "Got unexpected error")
	require.Equal(t, status, http.StatusOK)
	require.Equal(t, "[]", got, "Expected empty account list")
}
