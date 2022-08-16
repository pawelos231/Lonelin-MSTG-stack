package middleware

import (
	"BackendGo/pkg/models"
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var header = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjAsIk5hbWUiOiJYREREREREREREIiwiRW1haWwiOiJwYXdlbG9zQHBhd2Vsb3Nla2RzYWRhc2RzYSIsImV4cCI6MTY2MDc0NzQ2MH0.XrfTO1MCfNu-tg9IopTZ3Hm2x-Qd629hN3cOyn2QymM"
		var header2 = req.Header.Get("access-token")

		fmt.Println(header2)
		tk := &models.Token{}
		tkn, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
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

		context := context.WithValue(req.Context(), "user", tk)
		next.ServeHTTP(w, req.WithContext(context))
	})
}
