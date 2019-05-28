package template

import (
	"html/template"
	"net/http"
)

// IDs of templates
const (
	IndexTmplID int = 1
)

var (
	indexTmpl *template.Template
)

func init() {
	var err error
	indexTmpl, err = template.New("index.html").ParseFiles("tmpl/index.html")
	if err != nil {
		// handle error
	}
}

// ExecuteTemplate executes template of given ID (with optional data) to http writer
func ExecuteTemplate(templateID int, w *http.ResponseWriter, data interface{}) {
	switch templateID {
	case IndexTmplID:
		err := indexTmpl.ExecuteTemplate(*w, "index.html", data)
		if err != nil {
			// handle error
		}
	}
}
