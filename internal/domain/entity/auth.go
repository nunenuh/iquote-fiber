package entity

type Login struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}
