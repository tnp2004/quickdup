package utils

import "net/http"

func SetCookie(name, value, path string, maxAge int) http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		MaxAge:   maxAge,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
}
