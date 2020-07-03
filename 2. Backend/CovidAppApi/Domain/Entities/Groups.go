package entities

type GroupData struct {
	Edad               int     `json:"edad,omitempty"`
	Peso               float64 `json:"peso,omitempty"`
	Distrito           int     `json:"distrito,omitempty"`
	Tos                int     `json:"tos,omitempty"`
	Fiebre             int     `json:"fiebre,omitempty"`
	DificultadRespirar int     `json:"dificultadRespirar,omitempty"`
	PerdidaOlfato      int     `json:"perdidaOlfato,omitempty"`
	Enfermo            int     `json:"enfermo,omitempty"`
}

type ClusteredGroup struct {
	Name string
	Data []GroupData
}

type Clusters struct {
	Score    float64
	Clusters []ClusteredGroup
}

type GroupDataString struct {
	Edad               int     `json:"edad,omitempty"`
	Peso               float64 `json:"peso,omitempty"`
	Distrito           string  `json:"distrito,omitempty"`
	DificultadRespirar string  `json:"dificultadRespirar,omitempty"`
	PerdidaOlfato      string  `json:"perdidaOlfato,omitempty"`
}

type ClusteredGroupString struct {
	Name string
	Data []GroupDataString
}

type ClustersString struct {
	Score    float64
	Clusters []ClusteredGroupString
}
