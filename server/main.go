package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	consts "BackendGo/pkg/constants"
	Routes "BackendGo/pkg/routes"

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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(consts.MONGOURL))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)
	mime.AddExtensionType(".js", "application/javascript")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	router := mux.NewRouter()
	http.Handle("/user", Routes.UserHandlers(router, ctx, client))
	http.Handle("/post", Routes.PostsHandlers(router, ctx, client))
	http.Handle("/comments", Routes.CommentHandlers(router, ctx, client))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(corsObj, credentials)(router)))
}
