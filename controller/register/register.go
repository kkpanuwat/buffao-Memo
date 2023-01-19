package register

import (
	"kkpanuwat/buffaloMemo/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person represents a person with a name and age

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fname    string `json:"fname"`
}

func Register(c *gin.Context) {
	print(c)
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := orm.User{
		Username: json.Username,
		Password: json.Password,
		Fname:    json.Fname,
	}

	orm.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": json,
	})
}
