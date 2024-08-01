package route

import (
	"net/http"

	"maguro-alternative/go-go-go/route/todo"
)

func Routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.Todo)
}
