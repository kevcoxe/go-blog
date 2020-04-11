
package main


import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)



func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Homepage Endpoint Hit")
}


func handleRequests() {

  port := getEnv("PORT", "8080")
  fmt.Printf("Starting up the api on port %s\n", port)

  CreateTable()

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/", homePage)
  myRouter.HandleFunc("/posts", GetPosts).Methods("GET")
  myRouter.HandleFunc("/create/post", AddPost).Methods("POST")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}


func main() {
  handleRequests()
}

