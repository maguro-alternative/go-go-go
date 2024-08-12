package repository

import (
	"maguro-alternative/go-go-go/models"
)

func (t Repository) CreateTodo(content string) error {
	result := t.db.Create(&models.Todo{Content: content})
	return result.Error
}

func (t Repository) ReadAllTodo() ([]*models.Todo, error) {
	var todos []*models.Todo
	result := t.db.Find(&todos)
	return todos, result.Error
}

func (t Repository) ReadTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	result := t.db.First(&todo, id)
	return &todo, result.Error
}

func (t Repository) UpdateTodo(id int, content string) error {
	result := t.db.Where("id = ?", id).Updates(&models.Todo{Content: content})
	return result.Error
}

func (t Repository) DeleteTodo(id int) error {
	result := t.db.Delete(&models.Todo{}, id)
	return result.Error
}
