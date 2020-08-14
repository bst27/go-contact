package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")

		if name == "" || message == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad Request",
				"Message": "Name and message must not be empty",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Ok",
		})
	})

	_ = r.Run()
}
