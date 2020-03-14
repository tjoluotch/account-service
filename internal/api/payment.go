package api

import "net/http"

func (service *Service) PaymentHandler(resp http.ResponseWriter, req *http.Request) {
	logger := service.Logger
	logger.Info("payment")
}
