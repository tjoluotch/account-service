package api

import (
	"consul-service/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func (service *Service) PaymentHandler(resp http.ResponseWriter, req *http.Request) {
	logger := service.Logger
	logger.Info("payment handler: account service")
	//	TODO: implelment handler
	paymentModel := &models.Payment{}
	err := json.NewDecoder(req.Body).Decode(paymentModel)
	//	check for empty body - return bad request if EOF else 500 status code
	if err != nil {
		switch err {
		case io.EOF:
			logger.Errorw("empty request payload",
				"error", err)
			resp.WriteHeader(http.StatusBadRequest)
			return
		default:
			logger.Errorw("json decode error",
				"error", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//if payload := req.Body; payload == io.EOF {
	//
	//}
	//	decode request payload into Data type if err - return bad request
	//	generate payment id
	// comms with grpc
}
