package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	IsDone      bool   `json:"IsDone"`
}
