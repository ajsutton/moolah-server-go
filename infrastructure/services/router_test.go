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
		got, err := router.Call(http.MethodGet, "/")
		require.NoError(t, err)
		require.JSONEq(t, want, got)
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
		got, err := router.Call(http.MethodPost, "/")
		require.NoError(t, err)
		require.JSONEq(t, want, got)
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
