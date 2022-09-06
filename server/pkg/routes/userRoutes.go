package Routes

import (
	consts "BackendGo/pkg/constants"
	"BackendGo/pkg/controllers/UserController"
	"BackendGo/pkg/middleware"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserHandlers(r *mux.Router, ctx context.Context, client *mongo.Client) *mux.Router {
	collectionOfUsers := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_USERS)

	r.Use(CommonMiddlewareUser)
	r.HandleFunc("/createUser", UserController.Register(collectionOfUsers, ctx)).Methods("POST")
	r.HandleFunc("/loginUser", UserController.Login(collectionOfUsers, ctx)).Methods("POST")
	r.HandleFunc("/refreshToken", UserController.RefreshTokenHandler(collectionOfUsers, ctx)).Methods("POST", "GET")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.JwtVerify)
	s.HandleFunc("/userId", UserController.FetchUserById(collectionOfUsers, ctx)).Methods("GET", "OPTIONS")

	return r
}

func CommonMiddlewareUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET,PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, x-access-token")
		next.ServeHTTP(w, r)
	})
}
