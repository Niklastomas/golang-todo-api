package database

import (
	"log"

	"github.com/niklastomas/golang-todo-api/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoDB interface {
	Open() error
	Close() error
	CreateTodo(p *models.Todo) error
	GetTodo() ([]*models.Todo, error)
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Open() error {
	dsn := "host=localhost user=postgres password=postgres dbname=todoDB port=5432 sslmode=disable"

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	log.Println("Connected to Database!")

	d.db = pg
	return nil
}

func (d *DB) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (d *DB) Setup() error {
	err := d.db.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}
	seed(d.db)
	return nil
}

func seed(db *gorm.DB) {
	todos := []models.Todo{
		{Title: "Test1", Description: "TestTest", IsDone: false},
		{Title: "Test2", Description: "TestTest", IsDone: false},
		{Title: "Test3", Description: "TestTest", IsDone: true},
	}

	for _, todo := range todos {
		db.Create(&todo)
	}
}
