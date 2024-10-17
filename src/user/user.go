package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp.com/todo/database"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func CreateUser(c *gin.Context){
	db := database.GetDatabase()
	var user User
	if err := c.ShouldBindJSON(&user); err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Create(&user); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	db := database.GetDatabase()
	var users []User
	if err := db.Get(&users); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	db := database.GetDatabase()
	var user User
	userId, err := strconv.ParseUint(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	if err := db.Get(&user, userId); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
