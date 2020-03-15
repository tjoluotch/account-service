package api

import "net/http"

func (service *Service) PaymentHandler(resp http.ResponseWriter, req *http.Request) {
	logger := service.Logger
	logger.Info("payment")
	//	TODO: implelment handler
	//	check for empty body - return bad request
	//	decode request payload into Data type if
}
