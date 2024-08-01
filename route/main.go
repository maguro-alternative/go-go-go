package route

import (
	"net/http"

	"maguro-alternative/go-go-go/route/todo"
	"maguro-alternative/go-go-go/repository"
)

func Routes(repo repository.Repository) *http.ServeMux {
	mux := http.NewServeMux()
	todo := todo.NewTodoService(
		repo,
	)
	mux.HandleFunc("/", todo.Todo)
	return mux
}
