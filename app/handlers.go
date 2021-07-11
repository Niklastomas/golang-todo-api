package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/golang-todo-api/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to TODO api")
	}
}

func (a *App) CreateTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := models.Todo{}
		err := parse(w, r, &todo)
		if err != nil {
			log.Printf("Cannot parse todo body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.CreateTodo(&todo)
		if err != nil {
			log.Printf("Failed to create todo")
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, todo, http.StatusCreated)
	}
}

func (a *App) GetAllTodosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := a.DB.GetAllTodos()
		if err != nil {
			log.Printf("Cannot get todos, err=%v", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, todos, http.StatusOK)
	}
}

func (a *App) GetTodoByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		todo, err := a.DB.GetTodoById(id)
		if err != nil {
			log.Printf("Cannot find todo with id %s", id)
			sendResponse(w, r, nil, http.StatusNotFound)
			return
		}
		sendResponse(w, r, todo, http.StatusOK)
	}
}

func (a *App) DeleteTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		err := a.DB.DeleteTodo(id)
		if err != nil {
			log.Printf("Cannot delete todo with id %s", id)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}
		sendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *App) UpdateTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		todo := &models.Todo{}
		parse(w, r, todo)
		err := a.DB.UpdateTodo(id, todo)
		if err != nil {
			log.Println(err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}
		sendResponse(w, r, nil, http.StatusOK)
	}
}

func (a *App) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := parse(w, r, &user)
		if err != nil {
			log.Println(err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		newUser, err := a.DB.Register(user)
		if err != nil {
			log.Println(err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}
		sendResponse(w, r, newUser, http.StatusOK)
	}
}

func (a *App) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := parse(w, r, &user)
		if err != nil {
			log.Println(err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		user, err = a.DB.Login(user)
		if err != nil {
			log.Println(err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		sendResponse(w, r, user, http.StatusOK)

	}
}

func (a *App) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.DB.GetUsers()
		if err != nil {
			log.Fatalln(err)
			sendResponse(w, r, nil, http.StatusNotFound)
		}
		sendResponse(w, r, users, http.StatusOK)
	}
}
