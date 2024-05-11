package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
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

// CompressResponse сжимает ответ
func (a *Application) CompressResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(GzipWriter{Writer: gz, ResponseWriter: w}, r)
	})
}
