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

	CommonPostRouter := r.PathPrefix("/posts").Subrouter()

	CommonPostRouter.HandleFunc("/getdata", PostControllers.GetPostsData(CollectionOfPosts, ctx)).Methods("GET")
	CommonPostRouter.HandleFunc("/getSinglePost", PostControllers.GetPostByUniqueId(CollectionOfPosts, ctx)).Methods("GET")

	AuthenticateRouter := r.PathPrefix("/posts").Subrouter()

	//apply middleware
	AuthenticateRouter.Use(utils.CommonMiddleware)
	AuthenticateRouter.Use(middleware.JwtVerify)

	AuthenticateRouter.HandleFunc("/PostAPost", PostControllers.PostDataToDataBase(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	AuthenticateRouter.HandleFunc("/UpdatePost", PostControllers.UpdatePostFromDatabase(CollectionOfPosts, ctx)).Methods(http.MethodPut)

	AuthenticateRouter.HandleFunc("/DeletePost", PostControllers.DeletePostFromDatabase(CollectionOfPosts, ctx)).Methods(http.MethodDelete, http.MethodPost)

	AuthenticateRouter.HandleFunc("/FetchSpecificUserPosts", PostControllers.FetchUserSpecificPosts(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	AuthenticateRouter.HandleFunc("/DeleteAllPostsOfUser", PostControllers.DeleteAllPostsOfUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	return r
}
