package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"text/template"

	"github.com/ongoingio/site/app/model"
)

// List shows a list of all available examples.
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/layout.html", "templates/list.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "layout", model.List())
}
