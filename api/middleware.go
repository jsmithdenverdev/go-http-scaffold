package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
}

func useLogging(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			responseWriter := &loggingResponseWriter{
				ResponseWriter: rw,
			}

			next.ServeHTTP(responseWriter, r)

			logger.Printf("%d %s %s", responseWriter.status, r.Method, r.RequestURI)
		})
	}
}

func useTiming(logger *log.Logger, thresholdMs int64) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(rw, r)

			elapsed := time.Since(start).Milliseconds()

			if elapsed > thresholdMs {
				logger.Printf("%s %s took longer than threshold %dms ", r.Method, r.RequestURI, thresholdMs)
			}
		})
	}
}
