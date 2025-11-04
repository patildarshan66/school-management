package middlewares

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipResposeWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(w, r)
	})
}

type gzipResposeWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResposeWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
