package PostRoutes

import (
	PostControllers "BackendGo/pkg/controllers/PostController"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostsHandlers(r *mux.Router, col *mongo.Collection, ctx context.Context) *mux.Router {
	r.Use(CommonMiddleware)
	r.HandleFunc("/PostAPost", PostControllers.PostDataToDataBase(col, ctx)).Methods("POST")
	r.HandleFunc("/getdata", PostControllers.GetDataFromDatabase(col, ctx)).Methods("GET")
	r.HandleFunc("/getSinglePost", PostControllers.GetPostByUniqueId(col, ctx)).Methods("GET")
	return r

}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next.ServeHTTP(w, r)
	})
}
