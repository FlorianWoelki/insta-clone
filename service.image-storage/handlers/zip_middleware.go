package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GzipHandler struct
type GzipHandler struct {
}

// GzipMiddleware is a http middleware for handling gzip compression
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if given compression algorithm in the header is gzip
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gziped response
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}

		next.ServeHTTP(rw, r)
	})
}

// WrappedResponseWriter is a wrapper for including a gzip.Writer and http.ResponseWriter
type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter creates a new wrapped response writer for gziping
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gzipWriter := gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw: rw, gw: gzipWriter}
}

// Header returns the Header for the WrappedResponseWriter
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader writes a header with a given statuscode to the http.ResponseWriter
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

// Flush closes savely the gzip writer
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
