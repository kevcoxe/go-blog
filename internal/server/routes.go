package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/api/posts", s.GetPostsHandler)
	r.HandlerFunc(http.MethodPost, "/api/posts", s.AddPostHandlerJson)
	r.HandlerFunc(http.MethodDelete, "/api/posts/:id", s.DeletePostHandler)

	r.HandlerFunc(http.MethodGet, "/health", s.healthHandler)

	r.GET("/", s.ServeViteApp)
	r.ServeFiles("/assets/*filepath", http.Dir("./frontend/dist/assets"))

	return r
}

func (s *Server) ServeViteApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	fs.ServeHTTP(w, r)
}
