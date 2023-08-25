package domain

type Login struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
