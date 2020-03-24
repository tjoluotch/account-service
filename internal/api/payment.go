package api

import (
	"consul-service/internal/config"
	"consul-service/internal/models"
	"consul-service/internal/pb"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

var (
	errorJson = &ErrorResponse{Error: "grpc request failure"}
)

func (service *Service) PaymentHandler(resp http.ResponseWriter, req *http.Request) {
	logger := service.Logger
	logger.Info("payment handler: account service")

	paymentModel := &models.Payment{}
	logger.Info("Attempting to decode request payload to Payment struct")
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
	logger.Info("successfully decoded payload to Payment struct")

	//	generate payment id random 6 s.f. integer
	logger.Info("attempting to generate id")
	id, err := config.IdGenerator()
	if err != nil {
		logger.Errorw("error generating id",
			"error", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return

	}
	logger.Infow("ID generation successful",
		"id ", id)

	// comms with grpc via client
	client := *service.Client
	logger.Info("encoding rpc model")
	rpcModel := &pb.Payment{
		Amount:     paymentModel.Amount,
		SenderBank: paymentModel.SenderBank}

	_, err = client.SavePayment(context.Background(), rpcModel)
	if err != nil {
		logger.Errorw("rpc request failed",
			"error", err)
		resp.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(resp).Encode(errorJson)
		if err != nil {
			logger.Error("json encoding error", err)
		}
		return
	}
	resp.WriteHeader(http.StatusOK)
}
