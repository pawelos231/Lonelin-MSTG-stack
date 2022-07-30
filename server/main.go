package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func databases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
	w.WriteHeader(http.StatusOK)
	test := []string{}
	test = append(test, "Nice")
	test = append(test, "Nie")
	json.NewEncoder(w).Encode(test)
}

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

func PostDataToDataBase(client *mongo.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		var PostDetails PostInformiation
		json.NewDecoder(r.Body).Decode(&PostDetails)
		col := client.Database("MSTG_STACK").Collection("PostInfos")
		fmt.Println("Collection Type: ", reflect.TypeOf(col), "/n")

		result, insertErr := col.InsertOne(ctx, PostDetails)
		if insertErr != nil {
			fmt.Println("IntertOne Error", insertErr)
		} else {
			fmt.Println("insertOne() result type", reflect.TypeOf(result))
			fmt.Println("insertOne() api result type", result)
			newID := result.InsertedID
			fmt.Println("InsertedOne(), newID ", newID)
			fmt.Println("insertedOne(), newID type: ", reflect.TypeOf(newID))
		}
	}
}

func main() {
	//conmect to database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)

	mime.AddExtensionType(".js", "application/javascript")
	r := mux.NewRouter()

	r.HandleFunc("/homeland", PostDataToDataBase(client, ctx)).Methods("POST")
	http.Handle("/test", http.HandlerFunc(databases))
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
