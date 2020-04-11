package main

import (
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

var (
  database, _ = sql.Open("sqlite3", "./test.db")
)


func createTable() {
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS posts(post_id INTEGER PRIMARY KEY, title TEXT, content TEXT)")
  statement.Exec()
}


func addPostToDb(newPost Post) {
  fmt.Printf("adding post: %+v\n", newPost)
  statement, _ := database.Prepare("INSERT INTO posts (title, content) VALUES (?, ?)")
  statement.Exec(newPost.Title, newPost.Content)
}


func getPostsFromDb() Posts {
  posts := Posts{}
  rows, _ := database.Query("SELECT * FROM posts")

  var id int
  var title string
  var content string

  for rows.Next() {
    rows.Scan(&id, &title, &content)
    newPost := Post{Id:id, Title:title, Content:content}
    fmt.Printf("%+v\n", newPost)
    posts = append(posts, newPost)
  }

  return posts
}
