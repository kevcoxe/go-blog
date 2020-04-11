package main


type Post struct {
  Id int `json:"Id"`
  Title string `json:"Title"`
  Content string `json:"Content"`
}


type Picture struct {
  Id int `json:"Id"`
  PictureLocation string `json:"PictureURL"`
}


func CreatePost(id int, title string, content string) Post {
  post := Post{Id:id, Title:title, Content:content}

  return post
}

