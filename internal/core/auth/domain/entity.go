package domain

type Auth struct {
	ID          int    `json:"id,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	FullName    string `json:"full_name" validate:"required,min=2,max=100"`
	Phone       string `json:"phone"`
	IsSuperuser bool   `json:"is_superuser,omitempty"`
}

type CustomClaims struct {
	UserID      int
	Username    string
	Email       string
	IsSuperuser bool
}
