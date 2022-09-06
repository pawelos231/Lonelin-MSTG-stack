package auth

import (
	"net/http"
)

func SendRefreshToken(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jid",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	})
}
