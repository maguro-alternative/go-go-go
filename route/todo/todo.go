package todo

import (
	"encoding/json"
	"net/http"

	"maguro-alternative/go-go-go/repository"
	"maguro-alternative/go-go-go/route/todo/internal"
)

type TodoService struct {
	repo repository.Repository
}

func NewTodoService(
	repo repository.Repository,
) TodoService {
	return TodoService{
		repo: repo,
	}
}

func (t TodoService) Todo(w http.ResponseWriter, r *http.Request) {
	var todoJson internal.TodoJson
	err := json.NewDecoder(r.Body).Decode(&todoJson)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		readTodo(t.repo)
	case http.MethodPost:
		createTodo(t.repo, todoJson)
	case http.MethodPut:
		updateTodo(t.repo, todoJson)
	case http.MethodDelete:
		deleteTodo(t.repo, todoJson.ID)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Hello, World!"))
}

func createTodo(
	repo repository.Repository,
	todoJson internal.TodoJson,
) error {
	return repo.CreateTodo(todoJson.Name)
}

func readTodo(repo repository.Repository) ([]*internal.TodoJson, error) {
	var todos []*internal.TodoJson
	readTodos, err := repo.ReadTodo()
	if err != nil {
		return nil, err
	}
	for _, todo := range readTodos {
		todos = append(todos, &internal.TodoJson{
			ID:   todo.ID,
			Name: todo.Name,
		})
	}
	return todos, nil
}

func updateTodo(
	repo repository.Repository,
	todoJson internal.TodoJson,
) error {
	return repo.UpdateTodo(todoJson.ID, todoJson.Name)
}

func deleteTodo(repo repository.Repository, id int) error {
	return repo.DeleteTodo(id)
}
