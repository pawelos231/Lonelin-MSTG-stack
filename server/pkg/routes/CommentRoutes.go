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

	CollectionOfPosts := client.Database(consts.DATABASE_NAME).Collection(consts.COLLECTION_COMMENTS)

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(utils.CommonMiddleware)

	//later create subrouter for this endopoint
	r.HandleFunc("/GetAllCommentsOfGivenPosts", CommentsControllers.GetAllCommentsOfGivenPost(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	c := r.PathPrefix("/comments").Subrouter()

	//Apply middleware to controllers that need user aut
	c.Use(middleware.JwtVerify)

	c.HandleFunc("/CommentOnPostByUser", CommentsControllers.CommentOnPostByUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	c.HandleFunc("/LikeCommentOnPostByUser", CommentsControllers.LikeCommentOnPostByUser(CollectionOfPosts, ctx)).Methods(http.MethodPost)

	return r
}
