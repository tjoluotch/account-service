package api

import (
	"net/http"
)

const logMsg = "LOGGING MIDDLEWARE"

func (service *Service) LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		logger := service.Logger
		logger.Infow(logMsg, "request URI", req.RequestURI)
		logger.Infow(logMsg, "request method", req.Method)
		logger.Infow(logMsg, "IP", req.RemoteAddr)
		handler.ServeHTTP(resp, req)
	})
}
