package app

import (
	"github.com/gorilla/mux"
	"github.com/niklastomas/golang-todo-api/app/database"
)

type App struct {
	Router *mux.Router
	DB     database.DB
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	app.initRoutes()
	return app

}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/todo", a.CreateTodoHandler()).Methods("POST")
	a.Router.HandleFunc("/api/todo", a.GetAllTodosHandler()).Methods("GET")
	a.Router.HandleFunc("/api/todo/{id:[0-9]+}", a.UpdateTodoHandler()).Methods("PUT")
	a.Router.HandleFunc("/api/todo/{id:[0-9]+}", a.GetTodoByIdHandler()).Methods("GET")
	a.Router.HandleFunc("/api/todo/{id:[0-9]+}", a.DeleteTodoHandler()).Methods("DELETE")

	a.Router.HandleFunc("/api/register", a.RegisterHandler()).Methods("POST")
	a.Router.HandleFunc("/api/login", a.LoginHandler()).Methods("POST")

	a.Router.HandleFunc("/api/users", a.GetUsersHandler()).Methods("GET")
}
