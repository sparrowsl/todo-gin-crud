package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go Basics", Completed: true},
	{ID: "2", Item: "Read Some Books", Completed: false},
	{ID: "3", Item: "Use Neovim", Completed: false},
	{ID: "4", Item: "Learn Gin framework", Completed: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(ctx *gin.Context) {
		var newTodo todo

		if err := ctx.BindJSON(&newTodo); err != nil {
			return
		}

		todos = append(todos, newTodo)
		ctx.IndentedJSON(http.StatusCreated, newTodo)
	})

	router.GET("/todos/:id", func(ctx *gin.Context) {
		todoId := ctx.Param("id")

		for _, item := range todos {
			if item.ID == todoId {
				ctx.IndentedJSON(http.StatusOK, item)
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	router.PATCH("/todos/:id", func(ctx *gin.Context) {
		todoId := ctx.Param("id")

		for i, item := range todos {
			if item.ID == todoId {
				todos[i].Completed = !todos[i].Completed
				ctx.IndentedJSON(http.StatusOK, &todos[i])
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	router.DELETE("/todos/:id", func(ctx *gin.Context) {
		todoId := ctx.Param("id")

		for i, item := range todos {

			if item.ID == todoId {
				todos = append(todos[:i], todos[i+1:]...)
				ctx.IndentedJSON(http.StatusNoContent, item)
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	router.Run("localhost:3000")
}
