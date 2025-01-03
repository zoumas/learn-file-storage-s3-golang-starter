package main

import (
	"log"
	"net/http"
)

func middlewareLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCodeWriter := NewStatusCodeResponseWriter(w)
		next.ServeHTTP(statusCodeWriter, r)
		log.Println(r.Method, r.URL.Path, statusCodeWriter.statusCode)
	}
}

type StatusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewStatusCodeResponseWriter(w http.ResponseWriter) *StatusCodeResponseWriter {
	return &StatusCodeResponseWriter{ResponseWriter: w}
}

func (w *StatusCodeResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
