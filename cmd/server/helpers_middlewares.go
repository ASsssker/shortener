package main

import (
	"io"
	"net/http"
)

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

type GzipWriter struct {
	http.ResponseWriter
	Writer io.Writer // Writer будет использоваться для записи в ответ данных в сжатом виде
}

func (gw GzipWriter) Write(b []byte) (int, error) {
	return gw.Writer.Write(b)
}
