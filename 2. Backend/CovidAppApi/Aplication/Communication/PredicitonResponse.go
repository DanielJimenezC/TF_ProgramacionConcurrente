package communication

// Prediction Model
type PredictionResponse struct {
	Enfermo string    `json:"enfermo,omitempty"`
	Chance  float64 `json:"chance,omitempty"`
}
