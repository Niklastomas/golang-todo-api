package main

import (
	"log"
	"net/http"

	"github.com/niklastomas/golang-todo-api/app"
	"github.com/niklastomas/golang-todo-api/app/database"
)

func main() {
	app := app.New()
	app.DB = database.DB{}
	err := app.DB.Open()
	if err != nil {
		panic(err)
	}
	err = app.DB.Setup()
	if err != nil {
		log.Println(err)
	}

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
