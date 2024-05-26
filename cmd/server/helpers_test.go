package main

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptest"
	"shortener/internal/logger"

	"github.com/fatih/color"
)


// getTestApp создает тестовое приложеня
func getTestApp() *Application {
	app := &Application{
		InfoLog:  logger.CreateLogger("INFO", color.FgGreen, 0, nil),
		ErrorLog: logger.CreateLogger("ERROR", color.FgRed, 0, nil),
	}
	app.DatabaseDSN = ""
	app.FileStoragePath = ""
	app.connectToStorage()
	return app

}

// runHandler делает запрос к хэндлеру без запуска сервера
func runHandler(method string, url string, body string, f func(http.ResponseWriter, *http.Request)) *http.Response {
	r := httptest.NewRequest(method, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	f(w, r)
	return w.Result()
}

// testRequest делает запрос к хэндлеру через сервер
func testRequest(srv *httptest.Server, method string, body string) (*resty.Response, error) {
	req := resty.New().R()
	req.Method = method
	req.URL = srv.URL
	req.Body = body
	resp, err := req.Send()
	return resp, err
}