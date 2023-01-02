package main

import (
	"github.com/moolah-server-go/infrastructure/services"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListAccounts(t *testing.T) {
	application := NullApplication(Application{})
	router := application.router
	application.RegisterHandlers()

	status, got, err := router.Call(services.CallData{
		Method: http.MethodGet,
		Url:    "/api/accounts/",
	})
	require.Nil(t, err, "Got unexpected error")
	require.Equal(t, status, http.StatusOK)
	require.Equal(t, "[]", got, "Expected empty account list")
}
