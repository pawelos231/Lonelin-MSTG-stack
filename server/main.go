package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"BackendGo/pkg/routes/PostRoutes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//load env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	//TODO combine routes

	http.Handle("/", PostRoutes.PostsHandlers(r, col, ctx))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS()(r)))
}
