package route

import (
	"github.com/gin-gonic/gin"

	"maguro-alternative/go-go-go/repository"
	"maguro-alternative/go-go-go/route/todo"
)

func Routes(
	repo repository.Repository,
) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	t := todo.NewTodoService(repo)
	t.Todo(r)
	return r
}
