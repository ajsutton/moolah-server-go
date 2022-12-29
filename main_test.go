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

	got, err := router.Call(http.MethodPost, "/api/accounts/")
	require.Nil(t, err, "Got unexpected error")
	require.Empty(t, got, "Expected empty account list")
}
