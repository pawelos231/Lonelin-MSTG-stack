package middleware

import (
	"BackendGo/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//later place it in a header, i don't know why but cors keeps complaining when i try to set it

		//MOVE THIS ENTIRE THING TO SOME DIFFERENT FUNCTION CAUSE IT DUPLICATES IN USERCONTROLLER/REFRESHTOKEN

		if !utils.ValidateToken(w, req) {
			return
		}

		tokenValue := req.FormValue("q")
		tk := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
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
