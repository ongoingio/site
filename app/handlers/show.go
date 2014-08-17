package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/ongoingio/site/app/model"
)

// Show shows a single example by alias.
// TODO: Use path or sha instead, what happens there are multiple files with the same name in one folder?
func Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("DEBUG: getting example...")
	name, err := url.QueryUnescape(params.ByName("alias"))
	if err != nil {
		log.Fatal(err)
	}
	example, found := model.FindByName(name)
	if found != true {
		// TODO: Show 404
		log.Fatal("404 not found")
	}
	log.Printf("DEBUG: found example: %s", example)

	t, err := template.ParseFiles("templates/layout.html", "templates/show.html")
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, "layout", example)
}
