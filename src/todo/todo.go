package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp.com/todo/user"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	// ID uint64 `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CreatedBy uint `json:"createdBy" gorm:"unique;primaryKey;autoIncrement"`
	Assign uint `json:"assign" gorm:"unique;primaryKey;autoIncrement"`
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
	if !user.CheckUser(todo.Assign) {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "assigned user doesn't exist",
		})
		return
	}
	todos = append(todos, todo)
	c.JSON(http.StatusCreated,todo)
}

func GetTodo(c *gin.Context){
	c.JSON(http.StatusOK,todos)
}

func GetTodoById(c *gin.Context){
	todoId, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	for _,todo := range(todos){
		if uint64(todo.ID) == todoId{
			c.JSON(http.StatusOK,todo)
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"error": "no such todo exists"})
}

func ListUserTodo(c *gin.Context){
	userID, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	myTodo := []Todo{}
	for _,todo := range todos {
		if uint64(todo.Assign) == userID {
			myTodo = append(myTodo,todo)
		}
	}
	c.JSON(http.StatusOK,myTodo)
}

func UpdateTodo(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	var newTodo Todo
	if err:=c.ShouldBindJSON(&newTodo); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index:= range(todos){
		if uint64(todos[index].ID) == id{
			todos[index] = newTodo
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo updated successfully"})
}

func DeleteTodo(){

}