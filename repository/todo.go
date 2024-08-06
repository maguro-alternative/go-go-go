package repository

type todo struct {
	ID   int
	Name string
}

func (t Repository) CreateTodo(name string) error {
	result := t.db.Create(&todo{Name: name})
	return result.Error
}

func (t Repository) ReadAllTodo() ([]*todo, error) {
	rows, err := t.db.Find(&todo{}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*todo
	for rows.Next() {
		var todo todo
		err := rows.Scan(&todo.ID, &todo.Name)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	return todos, err
}

func (t Repository) ReadTodoByID(id int) (*todo, error) {
	var todo todo
	result := t.db.First(&todo, id)
	return &todo, result.Error
}

func (t Repository) UpdateTodo(id int, name string) error {
	result := t.db.Where("id = ?", id).Updates(&todo{Name: name})
	return result.Error
}

func (t Repository) DeleteTodo(id int) error {
	result := t.db.Delete(&todo{}, id)
	return result.Error
}
