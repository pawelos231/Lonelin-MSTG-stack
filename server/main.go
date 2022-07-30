package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
)

func databases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
	w.WriteHeader(http.StatusOK)
	test := []string{}
	test = append(test, "Spierdalaj")
	test = append(test, "Nie")
	json.NewEncoder(w).Encode(test)
}

type PostInformiation struct {
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	UserName  string `json:"userName"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	var PostDetails PostInformiation
	json.NewDecoder(r.Body).Decode(&PostDetails)
	fmt.Println(PostDetails)
}

func main() {

	mime.AddExtensionType(".js", "application/javascript")

	http.Handle("/test", http.HandlerFunc(databases))
	http.Handle("/homeland", http.HandlerFunc(HomeHandler))
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
