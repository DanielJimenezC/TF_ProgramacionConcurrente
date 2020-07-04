package communication

// UserResponse Model
type UserResponse struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}
