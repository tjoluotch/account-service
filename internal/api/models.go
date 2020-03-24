package api

import (
	"consul-service/internal/pb"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.SugaredLogger
	client *pb.AccountRoutesClient
}
