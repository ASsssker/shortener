package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	url    string
	body   string
	method string
}

type expect struct {
	status int
	body   string
}

type test struct {
	name  string
	value testData
	want  expect
}

func TestPostUrl(t *testing.T) {
	app := getTestApp()
	srv := httptest.NewServer(http.HandlerFunc(app.PostUrl))
	defer srv.Close()
	testCases := []test{
		{
			name: "Post method test #1",
			value: testData{
				url:    "/",
				body:   "{\"url\":\"https://go.dev/\"}",
				method: http.MethodPost,
			},
			want: expect{
				status: http.StatusCreated,
				body:   "",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := testRequest(srv, test.value.method, test.value.body)
			require.NoError(t, err)
			assert.Equal(t, test.want.status, res.StatusCode(), "Expected status code %d, got %d", test.want.status, res.StatusCode())
			require.NoError(t, err)
			assert.NotEmpty(t, res.Body(), "Response body is empty")
		})
	}
}

func TestIncorrectRequestGetUrl(t *testing.T) {
	app := getTestApp()
	srv := httptest.NewServer(http.HandlerFunc(app.GetUrl))
	defer srv.Close()
	testCases := []test{
		{
			name: "Invalid url test #1",
			value: testData{
				url:    "/novalid",
				method: http.MethodGet,
			},
			want: expect{
				status: http.StatusNotFound,
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := testRequest(srv, test.value.method, test.value.body)
			require.NoError(t, err)
			assert.Equal(t, test.want.status, res.StatusCode(), "Expected status code %d, got %d", test.want.status, res.StatusCode())
			require.NoError(t, err)
			assert.NotEmpty(t, res.Body(), "Response body is empty")
		})
	}
}

func TestCorrectRequestGetUrl(t *testing.T) {
	app := getTestApp()
	srv := httptest.NewServer(http.HandlerFunc(app.GetUrl))
	defer srv.Close()
	testCases := []test{
		{
			name: "Test #1",
			value: testData{
				url:    "/",
				body:   "https://go.dev/",
				method: http.MethodGet,
			},
			want: expect{
				status: http.StatusNotFound,
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := testRequest(srv, test.value.method, test.value.body)
			require.NoError(t, err)
			assert.Equal(t, test.want.status, res.StatusCode(), "Expected status code %d, got %d", test.want.status, res.StatusCode())
			require.NoError(t, err)
			assert.NotEmpty(t, res.Body(), "Response body is empty")
		})
	}
}
