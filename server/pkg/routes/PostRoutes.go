package Routes

import (
	PostControllers "BackendGo/pkg/controllers/PostController"
	"BackendGo/pkg/middleware"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostsHandlers(r *mux.Router, col *mongo.Collection, ctx context.Context) *mux.Router {
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CommonMiddleware)
	r.HandleFunc("/getdata", PostControllers.GetDataFromDatabase(col, ctx)).Methods("GET")
	r.HandleFunc("/getSinglePost", PostControllers.GetPostByUniqueId(col, ctx)).Methods("GET")

	s := r.PathPrefix("/").Subrouter()
	s.Use(CommonMiddleware)
	s.Use(middleware.JwtVerify)
	s.HandleFunc("/PostAPost", PostControllers.PostDataToDataBase(col, ctx)).Methods("POST")
	s.HandleFunc("/UpdatePost", PostControllers.UpdatePostFromDatabase(col, ctx)).Methods("PUT")
	s.HandleFunc("/DeletePost", PostControllers.DeletePostFromDatabase(col, ctx)).Methods(http.MethodDelete, http.MethodPost)

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
