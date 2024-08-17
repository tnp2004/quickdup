package authModels

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type Credentials struct {
	AccessToken  string
	RefreshToken string
}

type Authentication struct {
	UserID       string
	HashPassword string
	Password     string
}

type CredentialsResponse struct {
	AccessToken string `json:"accessToken"`
}

type AuthorizationCredentials struct {
	AccessToken  string
	RefreshToken string
}
