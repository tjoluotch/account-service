package api

import (
	"go.uber.org/zap"
	"net/http"
)

func HealthHandler(resp http.ResponseWriter, req *http.Request) {
	logger := zap.L().Sugar()
	logger.Infow("health check handler",
		"URI", req.RequestURI)
	_, _ = resp.Write([]byte("Health Check completed"))
	resp.WriteHeader(http.StatusOK)
}
