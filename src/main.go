package main

import (
	"github.com/gin-gonic/gin"
	"myapp.com/todo/todo"
	"myapp.com/todo/user"
	"myapp.com/todo/database"
	"fmt"
)

func main() {
	
	if err := database.CreateConnection("test.db",&user.User{},&todo.Todo{}); err != nil{
		panic(err.Error())
	} else {
		fmt.Println("Connection to database established")
	}

	server := gin.New()

	server.GET("/users",user.GetUsers)
	server.GET("/users/:id",user.GetUserById)
	server.POST("/users",user.CreateUser)

	server.POST("/todos",todo.CreateTodo)
	server.GET("/todos",todo.GetTodo)
	server.GET("/todos/:id",todo.GetTodoById)
	server.GET("/users/:id/todos",todo.ListUserTodo)
	server.PUT("/todos/:id",todo.UpdateTodo)

	server.Run("0.0.0.0:5000")
}