package api

import (
	"net/http"
	"time"
)

const logMsg = "LOGGING MIDDLEWARE"

func (service *Service) LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()
		logger := service.Logger
		logger.Infow(logMsg, "request URI", req.RequestURI,
			"request method", req.Method,
			"IP", req.RemoteAddr)
		method := req.Method + " "
		uri := req.RequestURI + " "
		ua := req.UserAgent() + " "
		handler.ServeHTTP(resp, req)
		completed := time.Since(start)
		logger.Info(method, uri, ua, completed)
	})
}
