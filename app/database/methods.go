package database

import (
	"log"

	"github.com/niklastomas/golang-todo-api/app/models"
)

func (d *DB) CreateTodo(todo *models.Todo) error {
	result := d.db.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DB) GetAllTodos() ([]*models.Todo, error) {
	var todos []*models.Todo
	result := d.db.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (d *DB) GetTodoById(id string) (*models.Todo, error) {
	var todo *models.Todo
	result := d.db.First(&todo, "id = ?", id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return todo, nil
}

func (d *DB) DeleteTodo(id string) error {
	result := d.db.Where("id = ?", id).Delete(&models.Todo{})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (d *DB) UpdateTodo(id string, todo *models.Todo) error {
	result := d.db.Model(&models.Todo{}).Where("id = ?", id).Updates(&todo)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
