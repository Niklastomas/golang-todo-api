package models

type User struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
}
