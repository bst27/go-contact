package main

import (
	"bytes"
	"encoding/json"
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

		err := sendWebhook("http://localhost:8081/webhook-receiver", name, message)
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

func sendWebhook(url string, name string, message string) error {
	body := make(map[string]string)
	body["name"] = name
	body["message"] = message

	jsonBody, err := json.MarshalIndent(body, "", "   ")
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(jsonBody))

	return err
}
