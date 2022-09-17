package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/madhushaw1012/Book-Management-System/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
