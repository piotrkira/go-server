package server

import (
	"net/http"

	"go-server/pkg/template"
	"go-server/services/article"
)

// IndexPage is representation of index page
type IndexPage struct {
	Title    string
	Articles []article.Article
}

func (s *Server) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data := &IndexPage{Title: "Index", Articles: []article.Article{article.GetLatestArticle()}}
		template.ExecuteTemplate(template.IndexTmplID, &w, data)
	}
}
