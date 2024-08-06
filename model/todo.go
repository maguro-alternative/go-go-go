package model

import "gorm.io/gorm"

type Todo struct {
	*gorm.Model
	Content string `json:"content"`
}
