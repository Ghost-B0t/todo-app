package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "myapp.com/todo/user"
	"myapp.com/todo/database"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	// ID uint64 `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CreatedBy uint `json:"createdBy"`
	Assign uint `json:"assign"`
	Status string `json:"status"`
}

var todos = []Todo{}

func CreateTodo(c *gin.Context){
	var todo Todo
	if err:= c.ShouldBindJSON(&todo); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := database.Create(&todo); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated,todo)
}

func GetTodo(c *gin.Context){
	var todos []Todo
	if err := database.Get(&todos); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context){
	var todo Todo
	todoId, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	if err := database.Get(&todo, todoId); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func ListUserTodo(c *gin.Context){
	userID, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	var userTodos []Todo
	// myTodo := []Todo{}
	if err := database.Get(&userTodos, &Todo{Assign: uint(userID)}); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,userTodos)
}

func UpdateTodo(c *gin.Context){
	todoId, err := strconv.ParseUint(c.Param("id"),10,64)
	// id := c.GetUint("id")
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	var newTodo Todo
	if err:=c.ShouldBindJSON(&newTodo); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todo Todo
	if err := database.Get(&todo, todoId); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	_ = database.Update(&todo,newTodo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(){

}