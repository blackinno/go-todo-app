package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
