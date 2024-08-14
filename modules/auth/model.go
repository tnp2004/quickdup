package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type LoginAuthentication struct {
	UserID       string
	HashPassword string
	Password     string
}
