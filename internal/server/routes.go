package server

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/old", s.IndexHandler)

	r.HandlerFunc(http.MethodGet, "/posts", s.GetPostsHandler)
	r.HandlerFunc(http.MethodPost, "/posts", s.AddPostHandlerJson)
	r.HandlerFunc(http.MethodDelete, "/posts/:id", s.DeletePostHandler)

	r.HandlerFunc(http.MethodGet, "/api/posts", s.GetPostsHandler)
	r.HandlerFunc(http.MethodPost, "/api/posts", s.AddPostHandlerJson)
	r.HandlerFunc(http.MethodDelete, "/api/posts/:id", s.DeletePostHandler)

	r.HandlerFunc(http.MethodPost, "/add-post", s.AddPostHandler)
	r.HandlerFunc(http.MethodPost, "/delete/:id", s.DeletePostHandler)
	r.HandlerFunc(http.MethodGet, "/health", s.healthHandler)

	r.GET("/", s.ServeViteApp)
	r.ServeFiles("/assets/*filepath", http.Dir("./frontend/dist/assets"))

	return r
}

func (s *Server) ServeViteApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	fs.ServeHTTP(w, r)
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageData := &PageData{
		Elements: []Element{},
	}

	pageData.Elements = append(pageData.Elements, &InputForm{
		Title:       "Create Post",
		Description: "Lets make a post",
		Elements: []FormElement{
			{
				Label: "Post Title",
				Type:  "text",
				Name:  "title",
				Value: "",
			},
			{
				Label: "Post Content",
				Type:  "textarea",
				Name:  "content",
				Value: "",
			},
		},
		Method: "Post",
		Action: "/add-post",
	})

	// get posts GetPosts
	posts, err := s.db.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, post := range posts {
		pageData.Elements = append(pageData.Elements, &Post{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		})
	}

	renderedPageData, err := pageData.Render()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var tmpl *template.Template = template.Must(template.ParseFiles("cmd/public/index.html"))
	if err := tmpl.Execute(w, renderedPageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
