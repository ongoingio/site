package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/ongoingio/site/app/examples"
)

// Show shows a single example by alias.
// TODO: Use path or sha instead, what happens there are multiple files with the same name in one folder?
func Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name, err := url.QueryUnescape(params.ByName("alias"))
	if err != nil {
		log.Fatal(err)
	}

	example, err := examples.FindByName(name)
	if err != nil {
		// TODO: Show 404
		log.Fatal("404 not found")
	}

	t, err := template.ParseFiles("templates/layout.html", "templates/show.html")
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, "layout", example)
}
