package entities

// Prediction Model
type Prediction struct {
	Edad               int     `json:"edad,omitempty"`
	Peso               float64 `json:"peso,omitempty"`
	Distrito           int     `json:"distrito,omitempty"`
	Tos                int     `json:"tos,omitempty"`
	Fiebre             int     `json:"fiebre,omitempty"`
	DificultadRespirar int     `json:"dificultadRespirar,omitempty"`
	PerdidaOlfato      int     `json:"perdidaOlfato,omitempty"`
	Enfermo            int     `json:"enfermo,omitempty"`
}
