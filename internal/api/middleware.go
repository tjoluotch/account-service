package api

import (
	"go.uber.org/zap"
	"net/http"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		logger := zap.L().Sugar()
		logger.Infow("request URI", req.RequestURI)
		logger.Infow("request method", req.Method)
		logger.Infow("source", req.RemoteAddr)
		handler.ServeHTTP(resp, req)
	})
}
