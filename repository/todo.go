package repository

import (
	"maguro-alternative/go-go-go/model"
)

func (t Repository) CreateTodo(content string) error {
	result := t.db.Create(&model.Todo{Content: content})
	return result.Error
}

func (t Repository) ReadAllTodo() ([]*model.Todo, error) {
	var todos []*model.Todo
	result := t.db.Find(&todos)
	return todos, result.Error
}

func (t Repository) ReadTodoByID(id int) (*model.Todo, error) {
	var todo model.Todo
	result := t.db.First(&todo, id)
	return &todo, result.Error
}

func (t Repository) UpdateTodo(id int, content string) error {
	result := t.db.Where("id = ?", id).Updates(&model.Todo{Content: content})
	return result.Error
}

func (t Repository) DeleteTodo(id int) error {
	result := t.db.Delete(&model.Todo{}, id)
	return result.Error
}
