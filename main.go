package main

import (
	Register "kkpanuwat/buffaloMemo/controller/register"
	"kkpanuwat/buffaloMemo/orm"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    "Buffalo Memo Web Application",
		})
	})
	r.POST("/register", Register.Register)
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
