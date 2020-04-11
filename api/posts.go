package main

import (
  "net/http"
  "encoding/json"
)


type Posts []Post


func GetPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  posts := GetPostsFromDb()
  json.NewEncoder(w).Encode(posts)
}


func AddPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var newPost Post
  _ = json.NewDecoder(r.Body).Decode(&newPost)
  AddPostToDb(newPost)
  json.NewEncoder(w).Encode(newPost)
}

