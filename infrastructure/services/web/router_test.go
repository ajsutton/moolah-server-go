package web

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		var want = "{\"a\": \"b\"}"
		require.NoError(t, router.Start(0))
		router.Get("/", func(request Request) (any, error) {
			output := map[string]string{"a": "b"}
			return output, nil
		})
		statusCode, got, err := router.Call(CallData{Method: http.MethodGet, Url: "/"})
		require.NoError(t, err)
		require.JSONEq(t, want, got)
		require.Equal(t, statusCode, http.StatusOK)
	})
}

func Test_Get_ReturnsError(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		require.NoError(t, router.Start(0))
		router.Get("/", func(request Request) (any, error) {
			return nil, BadRequest("Woah!")
		})
		statusCode, got, err := router.Call(CallData{Method: http.MethodGet, Url: "/"})
		require.Equal(t, statusCode, http.StatusBadRequest)
		require.NoError(t, err)
		require.Equal(t, "\"Woah!\"", got)
	})
}

func Test_Get_ReturnsUnexpectedError(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		require.NoError(t, router.Start(0))
		router.Get("/", func(request Request) (any, error) {
			return nil, CallError("Woah!")
		})
		statusCode, got, err := router.Call(CallData{Method: http.MethodGet, Url: "/"})
		require.Equal(t, statusCode, http.StatusInternalServerError)
		require.NoError(t, err)
		require.Equal(t, "\"Woah!\"", got)
	})
}

func Test_Post(t *testing.T) {
	runRouterTests(t, func(t *testing.T, router Router) {
		var want = "{\"a\": \"b\"}"
		sentData := TestData{"some value"}
		var receivedData TestData
		require.NoError(t, router.Start(0))
		router.Post("/", func(request Request) (any, error) {
			output := map[string]string{"a": "b"}

			err := request.BodyJson(&receivedData)
			if err != nil {
				return nil, BadRequest(err.Error())
			}
			return output, nil
		})
		statusCode, got, err := router.Call(CallData{
			Method: http.MethodPost,
			Url:    "/",
			Data:   sentData,
		})
		require.NoError(t, err)
		require.JSONEq(t, want, got)
		require.Equal(t, statusCode, http.StatusOK)
		require.Equal(t, sentData, receivedData)
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

type TestData struct {
	Value string `json:"value"`
}
