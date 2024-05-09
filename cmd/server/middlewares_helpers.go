package main

import "net/http"

type (
	ResponseInfo struct {
		size   int
		status int
	}

	loggingResponseWriter struct {
		ResponseInfo
		http.ResponseWriter
	}
)

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	lrw.status = status
	lrw.ResponseWriter.WriteHeader(status)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}
