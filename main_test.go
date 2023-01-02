package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListAccounts(t *testing.T) {
	application := NullApplication(Application{})
	router := application.router
	application.RegisterHandlers()

	status, got, err := router.Call(http.MethodGet, "/api/accounts/")
	require.Nil(t, err, "Got unexpected error")
	require.Equal(t, status, http.StatusOK)
	require.Equal(t, "[]", got, "Expected empty account list")
}
