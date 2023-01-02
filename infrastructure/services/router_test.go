package services

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		var want = "{\"a\": \"b\"}"
		require.NoError(t, router.Start(0))
		router.Get("/", func() (any, error) {
			output := map[string]string{"a": "b"}
			return output, nil
		})
		statusCode, got, err := router.Call(http.MethodGet, "/")
		require.NoError(t, err)
		require.JSONEq(t, want, got)
		require.Equal(t, statusCode, http.StatusOK)
	})
}

func Test_Get_ReturnsError(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		require.NoError(t, router.Start(0))
		router.Get("/", func() (any, error) {
			return nil, BadRequest("Woah!")
		})
		statusCode, got, err := router.Call(http.MethodGet, "/")
		require.Equal(t, statusCode, http.StatusBadRequest)
		require.NoError(t, err)
		require.Equal(t, "\"Woah!\"", got)
	})
}

func Test_Post(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		var want = "{\"a\": \"b\"}"
		require.NoError(t, router.Start(0))
		router.Post("/", func() (any, error) {
			output := map[string]string{"a": "b"}
			return output, nil
		})
		statusCode, got, err := router.Call(http.MethodPost, "/")
		require.NoError(t, err)
		require.JSONEq(t, want, got)
		require.Equal(t, statusCode, http.StatusOK)
	})
}

func runRouterTests(t *testing.T, f func(t *testing.T, router Router)) {
	var tests = []struct {
		name   string
		router Router
	}{
		{"Gin", NewRouter()},
		{"Null", NullRouter()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f(t, test.router)
		})
	}
}
