package authException

type EmailOrPasswordInvalid struct{}

func (e *EmailOrPasswordInvalid) Error() string {
	return "email or password invalid"
}
