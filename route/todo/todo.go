package todo

import (
	"net/http"
)

func Todo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
