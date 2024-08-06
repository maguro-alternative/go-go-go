package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

func (t TodoService) Todo(r *gin.Engine) {
	var todoJson internal.TodoJson
	r.POST("/todo", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&todoJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := createTodo(t.repo, todoJson); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Todo created"})
	})
	r.GET("/todo", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			todo, err := readTodoByID(t.repo, id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"todo": todo})
			return
		}
		todos, err := readAllTodo(t.repo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})
	r.PUT("/todo", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&todoJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := updateTodo(t.repo, todoJson); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
	})
	r.DELETE("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		if err := deleteTodo(t.repo, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
	})
}

func createTodo(
	repo repository.Repository,
	todoJson internal.TodoJson,
) error {
	return repo.CreateTodo(todoJson.Name)
}

func readAllTodo(repo repository.Repository) ([]*internal.TodoJson, error) {
	var todos []*internal.TodoJson
	readTodos, err := repo.ReadAllTodo()
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

func readTodoByID(
	repo repository.Repository,
	id int,
) (*internal.TodoJson, error) {
	todo, err := repo.ReadTodoByID(id)
	if err != nil {
		return nil, err
	}
	return &internal.TodoJson{
		ID:   todo.ID,
		Name: todo.Name,
	}, nil
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
