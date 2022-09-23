package CommentsControllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var Comment any
		json.NewDecoder(req.Body).Decode(&Comment)
		fmt.Println(Comment)
		json.NewEncoder(w).Encode("SIEMA")
	}
}
func LikeCommentOnPostByUser(col *mongo.Collection, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
