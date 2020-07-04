package services

import (
	entity "../../domain/entities"
)

// IPredictionService Interface
type IPredictionService interface {
	Predict(prediction entity.Prediction) (bool, float64, error)
}
