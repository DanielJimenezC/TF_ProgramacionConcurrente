package entities

// Prediction Model
type Prediction struct {
	Edad               string `json:"edad,omitempty"`
	Peso               string `json:"peso,omitempty"`
	Distrito           string `json:"distrito,omitempty"`
	Tos                string `json:"tos,omitempty"`
	Fiebre             string `json:"fiebre,omitempty"`
	DificultadRespirar string `json:"dificultadRespirar,omitempty"`
	PerdidaOlfato      string `json:"perdidaOlfato,omitempty"`
	Enfermo            string `json:"enfermo,omitempty"`
}
