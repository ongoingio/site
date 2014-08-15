package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "html/template"
)

// Index shows a simple home page.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "layout", nil)
}

// Examples shows an index of all available examples.
func Examples(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Examples!\n")
}

// Example shows a single example by name.
func Example(w http.ResponseWriter, r *http.Request, routes httprouter.Params) {
	fmt.Fprint(w, "Example " + routes.ByName("name") + "!\n")
}

func main() {
    router := httprouter.New()

    router.GET("/", Index)
    router.GET("/examples", Examples)
    router.GET("/examples/:name", Example)

    log.Fatal(http.ListenAndServe(":8080", router))
}