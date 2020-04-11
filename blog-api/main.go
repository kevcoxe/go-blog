
package main


import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)


func apiHomePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Homepage Endpoint Hit")
}


func handleRequests() {

  port := getEnv("PORT", "8080")
  fmt.Printf("Starting up the api on port %s\n", port)

  createTable()

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/", apiHomePage)
  myRouter.HandleFunc("/api/posts", getPosts).Methods("GET")
  myRouter.HandleFunc("/api/posts", addPost).Methods("POST")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}


func main() {
  handleRequests()
}

