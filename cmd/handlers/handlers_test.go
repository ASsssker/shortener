package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
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
	tests := []test{
		{
			name: "Get method test #1",
			value: testData{
				url:    "/",
				body:   "https://go.dev/",
				method: http.MethodGet,
			},
			want: expect{
				status: http.StatusMethodNotAllowed,
				body:   "Method Not Allowed",
			},
		},
		{
			name: "Post method test #1",
			value: testData{
				url:    "/",
				body:   "https://go.dev/",
				method: http.MethodPost,
			},
			want: expect{
				status: http.StatusCreated,
				body:   "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := request(test.value.method, test.value.url, test.value.body, PostUrl)
			defer res.Body.Close()
			assert.Equal(t, test.want.status, res.StatusCode, "Expected status code %d, got %d", test.want.status, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.NotEmpty(t, resBody, "Response body is empty")
		})
	}
}

func TestIncorrectRequestGetUrl(t *testing.T) {
	tests := []test{
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
		{
			name: "Post method test #1",
			value: testData{
				url:    "/novalid",
				method: http.MethodPost,
			},
			want: expect{
				status: http.StatusMethodNotAllowed,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := request(test.value.method, test.value.url, test.value.body, GetUrl)
			defer res.Body.Close()
			assert.Equal(t, test.want.status, res.StatusCode, "Expected status code %d, got %d", test.want.status, res.StatusCode)
			_, err := io.ReadAll(res.Body)
			require.NoError(t, err)
		})
	}
}

func TestCorrectRequestGetUrl(t *testing.T) {
	tests := []test{
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			postUrl := request(http.MethodPost, test.value.url, test.value.body, PostUrl)
			defer postUrl.Body.Close()
			shortUrl, err := io.ReadAll(postUrl.Body)
			require.NoError(t, err)
			test.value.url += string(shortUrl)
			res := request(test.value.method, test.value.url, test.value.body, GetUrl)
			defer res.Body.Close()
			assert.Equal(t, test.want.status, res.StatusCode)

		})
	}
}
