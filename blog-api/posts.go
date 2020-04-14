package main

import (
	"encoding/json"
	"net/http"
)

type Posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := getPostsFromDb()
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPost Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)
	addPostToDb(newPost)
	json.NewEncoder(w).Encode(newPost)
}
