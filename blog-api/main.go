package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var PORT string = getEnv("PORT", "8080")

func apiHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	fmt.Printf("Starting up the api on port %s\n", PORT)

	createTable()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", apiHomePage)
	myRouter.HandleFunc("/api/posts", getPosts).Methods("GET")
	myRouter.HandleFunc("/api/posts", addPost).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), myRouter))
}

func main() {
	handleRequests()
}
