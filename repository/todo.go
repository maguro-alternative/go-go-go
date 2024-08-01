package repository

type todo struct {
	ID   int
	Name string
}

func (t Repository) CreateTodo(name string) error {
	_, err := t.db.Exec("INSERT INTO todos (name) VALUES ($1)", name)
	return err
}

func (t Repository) ReadTodo() ([]*todo, error) {
	rows, err := t.db.Query("SELECT * FROM todos")
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

func (t Repository) UpdateTodo(id int, name string) error {
	_, err := t.db.Exec("UPDATE todos SET name = $1 WHERE id = $2", name, id)
	return err
}

func (t Repository) DeleteTodo(id int) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
