package main

import (
	"github.com/moolah-server-go/infrastructure/services"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListAccounts(t *testing.T) {
	router := services.NullRouter()
	accounts := services.NullAccounts()
	application := NullApplication(Application{accounts: accounts, router: router})
	application.RegisterHandlers()

	status, got, err := router.Call(http.MethodGet, "/api/accounts/")
	require.Nil(t, err, "Got unexpected error")
	require.Equal(t, status, http.StatusOK)
	require.Equal(t, "[]", got, "Expected empty account list")
}
