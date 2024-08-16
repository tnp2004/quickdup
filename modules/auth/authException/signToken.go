package authException

type SignAccessToken struct{}

func (e *SignAccessToken) Error() string {
	return "sign token failed"
}

type SignRefreshToken struct{}

func (e *SignRefreshToken) Error() string {
	return "sign token failed"
}

type UnexpectedSigningMethod struct{}

func (e *UnexpectedSigningMethod) Error() string {
	return "unexpected signing method"
}
