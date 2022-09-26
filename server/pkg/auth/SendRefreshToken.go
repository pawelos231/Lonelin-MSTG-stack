package auth

import (
	"net/http"
)

func SendRefreshToken(w http.ResponseWriter, req *http.Request, token string) {
	cookie := &http.Cookie{
		Name:     "jid",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}
