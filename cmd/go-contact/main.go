package main

import (
	"fmt"
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

		err := sendWebhook(name, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error":   "Internal Server Error",
				"Message": "Failed to process request",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Ok",
		})
	})

	_ = r.Run()
}

func sendWebhook(name string, message string) error {
	fmt.Println(name, message) // TODO: Implement
	return nil
}
