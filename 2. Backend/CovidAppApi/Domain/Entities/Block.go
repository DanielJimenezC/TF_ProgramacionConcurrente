package entities

// Block struct
type Block struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Dni      string `json:"dni,omitempty"`
	Telefono string `json:"telefono,omitempty"`
	Previous string `json:"previous,omitempty"`
	Hash     string `json:"hash,omitempty"`
}
