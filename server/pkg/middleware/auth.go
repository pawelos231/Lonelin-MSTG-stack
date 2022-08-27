package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("test")
		tokenValue := req.FormValue("q")
		tk := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			println("blad")
			return

		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		Context := context.WithValue(req.Context(), "user", tk)
		next.ServeHTTP(w, req.WithContext(Context))
	})
}
