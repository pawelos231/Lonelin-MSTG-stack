package Routes

import (
	consts "BackendGo/pkg/constants"
	PostControllers "BackendGo/pkg/controllers/PostController"
	"BackendGo/pkg/middleware"
	"BackendGo/pkg/utils"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostsHandlers(r *mux.Router, ctx context.Context, client *mongo.Client) *mux.Router {

	CollectionOfPosts := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_POSTS)

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(utils.CommonMiddleware)
	r.HandleFunc("/getdata", PostControllers.GetPostsData(CollectionOfPosts, ctx)).Methods("GET")
	r.HandleFunc("/getSinglePost", PostControllers.GetPostByUniqueId(CollectionOfPosts, ctx)).Methods("GET")

	s := r.PathPrefix("/").Subrouter()

	//apply middleware
	s.Use(utils.CommonMiddleware)
	s.Use(middleware.JwtVerify)

	s.HandleFunc("/PostAPost", PostControllers.PostDataToDataBase(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	s.HandleFunc("/UpdatePost", PostControllers.UpdatePostFromDatabase(CollectionOfPosts, ctx)).Methods(http.MethodPut)

	s.HandleFunc("/DeletePost", PostControllers.DeletePostFromDatabase(CollectionOfPosts, ctx)).Methods(http.MethodDelete, http.MethodPost)

	s.HandleFunc("/FetchSpecificUserPosts", PostControllers.FetchUserSpecificPosts(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	s.HandleFunc("/DeleteAllPostsOfUser", PostControllers.DeleteAllPostsOfUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	return r
}
