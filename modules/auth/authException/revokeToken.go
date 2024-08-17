package authException

type RevokeToken struct{}

func (e *RevokeToken) Error() string {
	return "revoke token failed"
}
