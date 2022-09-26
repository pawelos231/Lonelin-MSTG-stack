package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//later place it in a header, i don't know why but cors keeps complaining when i try to set it

		//MOVE THIS ENTIRE THING TO SOME DIFFERENT FUNCTION CAUSE IT DUPLICATES IN USERCONTROLLER/REFRESHTOKEN
		for _, c := range req.Cookies() {
			fmt.Println(c, "chuj")
		}
		tokenCookie2, errCookie := req.Cookie("jid")
		if errCookie != nil {
			fmt.Println(errCookie)
			return
		}

		value := tokenCookie2.Value

		tkClaims := jwt.MapClaims{}
		refreshToken, errParsed := jwt.ParseWithClaims(value, tkClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
		})
		//Pass it to utils later beacuse it duplicates
		if errParsed != nil || refreshToken == nil {
			if errParsed == jwt.ErrSignatureInvalid {
				json.NewEncoder(w).Encode("invalid token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("zły token")
			println("blad")
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
