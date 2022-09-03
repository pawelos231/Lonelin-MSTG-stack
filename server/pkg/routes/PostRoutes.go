package Routes

import (
	consts "BackendGo/pkg/constants"
	PostControllers "BackendGo/pkg/controllers/PostController"
	"BackendGo/pkg/middleware"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostsHandlers(r *mux.Router, ctx context.Context, client *mongo.Client) *mux.Router {

	CollectionOfPosts := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_POSTS)
	
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CommonMiddleware)
	r.HandleFunc("/getdata", PostControllers.GetDataFromDatabase(CollectionOfPosts, ctx)).Methods("GET")
	r.HandleFunc("/getSinglePost", PostControllers.GetPostByUniqueId(CollectionOfPosts, ctx)).Methods("GET")

	s := r.PathPrefix("/").Subrouter()
	s.Use(CommonMiddleware)
	s.Use(middleware.JwtVerify)
	s.HandleFunc("/PostAPost", PostControllers.PostDataToDataBase(CollectionOfPosts, ctx)).Methods("POST")
	s.HandleFunc("/UpdatePost", PostControllers.UpdatePostFromDatabase(CollectionOfPosts, ctx)).Methods("PUT")
	s.HandleFunc("/DeletePost", PostControllers.DeletePostFromDatabase(CollectionOfPosts, ctx)).Methods(http.MethodDelete, http.MethodPost)
	s.HandleFunc("/FetchSpecificUserPosts", PostControllers.FetchUserSpecificPosts(CollectionOfPosts, ctx)).Methods(http.MethodPost)

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
