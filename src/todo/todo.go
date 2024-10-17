package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp.com/todo/database"
)

type Todo struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	CreatedBy uint `json:"createdBy" binding:"required"`
	Assign uint `json:"assign"`
	Status string `json:"status" form:"status,default=Todo"`
}

func CreateTodo(c *gin.Context){
	db := database.GetDatabase()
	var todo Todo
	if err:= c.ShouldBindJSON(&todo); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Create(&todo); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated,todo)
}

func GetTodo(c *gin.Context){
	db := database.GetDatabase()
	var todos []Todo
	if err := db.Get(&todos); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context){
	db := database.GetDatabase()
	var todo Todo
	todoId, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	if err := db.Get(&todo, todoId); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func ListUserTodo(c *gin.Context){
	db := database.GetDatabase()
	userID, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	var userTodos []Todo
	if err := db.Get(&userTodos, &Todo{Assign: uint(userID)}); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,userTodos)
}

func UpdateTodo(c *gin.Context){
	db := database.GetDatabase()
	todoId, err := strconv.ParseUint(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	var newTodo Todo
	if err:=c.ShouldBindJSON(&newTodo); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todo Todo
	if err := db.Get(&todo, todoId); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	_ = db.Update(&todo,newTodo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(){

}