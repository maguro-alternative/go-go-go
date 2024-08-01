package repository

type Todo struct {
	ID   int
	Name string
}

func (t Repository) CreateTodo(todo Todo) error {
	_, err := t.db.Exec("INSERT INTO todos (name) VALUES ($1)", todo.Name)
	return err
}

func (t Repository) ReadTodo() ([]*Todo, error) {
	rows, err := t.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Name)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	return todos, err
}

func (t Repository) UpdateTodo(todo Todo) error {
	_, err := t.db.Exec("UPDATE todos SET name = $1 WHERE id = $2", todo.Name, todo.ID)
	return err
}

func (t Repository) DeleteTodo(id int) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
