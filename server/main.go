package main

import (
	"BackendGo/pkg/controllers"
	"context"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//conmect to database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)
	col := client.Database("MSTG_STACK").Collection("PostInfos")
	mime.AddExtensionType(".js", "application/javascript")
	r := mux.NewRouter()

	r.HandleFunc("/PostAPost", controllers.PostDataToDataBase(col, ctx)).Methods("POST")
	r.HandleFunc("/getdata", controllers.GetDataFromDatabase(col, ctx)).Methods("GET")
	r.HandleFunc("/getSinglePost", controllers.GetPostByUniqueId(col, ctx)).Methods("GET")

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
