package main

import (
	"net/http"
	"time"
)

// RequestsInfo логирует информацию о запросе
func (a *Application) RequestsInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			ResponseInfo:   ResponseInfo{},
		}
		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		a.InfoLog.Printf("URI: %s\tMethod: %s\tTime: %v\tSize: %d\tStatus: %d", r.RequestURI, r.Method, duration, lrw.size, lrw.status)
	})
}
