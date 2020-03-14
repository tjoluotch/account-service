package api

import (
	"net/http"
)

func (service *Service) HealthHandler(resp http.ResponseWriter, req *http.Request) {
	logger := service.Logger
	logger.Infow("health check handler",
		"URI", req.RequestURI)
	_, _ = resp.Write([]byte("Health Check completed"))
	resp.WriteHeader(http.StatusOK)
}
