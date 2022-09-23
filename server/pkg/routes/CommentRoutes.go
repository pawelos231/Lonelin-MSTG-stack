package Routes

import (
	consts "BackendGo/pkg/constants"
	CommentsControllers "BackendGo/pkg/controllers/CommentsController"
	"BackendGo/pkg/middleware"
	"BackendGo/pkg/utils"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentHandlers(r *mux.Router, ctx context.Context, client *mongo.Client) *mux.Router {

	CollectionOfPosts := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_POSTS)

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(utils.CommonMiddleware)

	c := r.PathPrefix("/comments").Subrouter()
	c.Use(middleware.JwtVerify)

	c.HandleFunc("/CommentOnPostByUser", CommentsControllers.CommentOnPostByUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	c.HandleFunc("/LikeCommentOnPostByUser", CommentsControllers.LikeCommentOnPostByUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	return r
}
