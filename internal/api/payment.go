package api

import (
	"consul-service/internal/models"
	"crypto/rand"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
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
	generator := rand.Reader
	max := big.NewInt(1000000)
	min := big.NewInt(100000)
	diff := max.Sub(max, min)
	res, err := rand.Int(generator, diff)
	if err != nil {
		logger.Errorw("error generating id",
			"error", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	res = res.Add(res, min)
	logger.Info("id ", res.String())

	// comms with grpc
}
