package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"time"

	"BackendGo/pkg/routes/PostRoutes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//conmect to database
	const (
		MONGOURL         = "mongodb://localhost:27017"
		DATABASE_NAME    = "MSTG_STACK"
		COLLECTION_POSTS = "PostInfos"
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOURL))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)
	col := client.Database(DATABASE_NAME).Collection(COLLECTION_POSTS)
	mime.AddExtensionType(".js", "application/javascript")
	r := mux.NewRouter()

	http.Handle("/", PostRoutes.PostsHandlers(r, col, ctx))

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
