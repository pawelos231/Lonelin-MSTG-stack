package userRoutes

import (
	consts "BackendGo/pkg/constants"
	"BackendGo/pkg/controllers/UserController"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserHandlers(r *mux.Router, ctx context.Context, client *mongo.Client) *mux.Router {
	col := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_USERS)
	r.Use(CommonMiddleware)
	r.HandleFunc("/createUser", UserController.Register(col, ctx)).Methods("POST")
	return r
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
