package entities

// GroupData struct
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

// ClusteredGroup struct
type ClusteredGroup struct {
	Name string
	Data []GroupData
}

// Clusters struct
type Clusters struct {
	Score    float64
	Clusters []ClusteredGroup
}

// GroupDataString struct
type GroupDataString struct {
	Edad               int     `json:"edad,omitempty"`
	Peso               float64 `json:"peso,omitempty"`
	Distrito           string  `json:"distrito,omitempty"`
	DificultadRespirar string  `json:"dificultadRespirar,omitempty"`
	PerdidaOlfato      string  `json:"perdidaOlfato,omitempty"`
}

// ClusteredGroupString struct
type ClusteredGroupString struct {
	Name string
	Data []GroupDataString
}

// ClustersString struct
type ClustersString struct {
	Score    float64
	Clusters []ClusteredGroupString
}
