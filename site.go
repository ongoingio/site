package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/ongoingio/site/app/database"
	"github.com/ongoingio/site/app/examples"
	"github.com/ongoingio/site/app/handlers"
)

/*
// Init loads and decodes the config file.
func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config = Config{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}
*/

func main() {
	database.Connect()

	examples.Register()
	examples.Sync()

	router := httprouter.New()
	router.GET("/", handlers.List)
	router.GET("/examples/:alias", handlers.Show)
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
