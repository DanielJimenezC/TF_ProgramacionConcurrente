package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	interfaces "../../aplication/interfaces"
	service "../../aplication/services"
	header "../../config"
	model "../../domain/entities"
	servError "../errors"
)

type predictcontroller struct{}

var (
	predictServ interfaces.IPredictionService = service.PredictionService()
)

// IPredictionController interface
type IPredictionController interface {
	Predict(response http.ResponseWriter, request *http.Request)
}

// PredictionController Implementation
func PredictionController() IPredictionController {
	return &predictcontroller{}
}

// Predict if sick
func (*predictcontroller) Predict(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}

	var predict model.Prediction
	err := json.NewDecoder(request.Body).Decode(&predict)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error decode Json users"})
	}

	result, _ := predictServ.Predict(predict)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err.Error()})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(strconv.FormatBool(result))
	}
}
