package api

import (
	"consul-service/internal/pb"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.SugaredLogger
	Client *pb.AccountRoutesClient
}

type ErrorResponse struct {
	Error string `json:"error"`
}
