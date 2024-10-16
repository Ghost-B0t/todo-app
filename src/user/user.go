package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp.com/todo/database"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

var users = []User{}

func CreateUser(c *gin.Context){
	var user User
	if err := c.ShouldBindJSON(&user); err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	_ = database.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []User
	_ = database.Get(&users)
	c.JSON(http.StatusOK, users)
}

func CheckUser(uid uint) bool{
	for _, user := range(users) {
		if user.ID == uid{
			return true
		}
	}
	return false
}
