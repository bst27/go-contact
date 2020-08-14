package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	url := flag.String("url", "", "URL to receive form submissions (webhook target)")
	nameKey := flag.String("nameKey", "name", "The webhook target will receive the name value of a form submission with this key in the json data")
	messageKey := flag.String("messageKey", "message", "The webhook target will receive the message value of a form submission with this key in the json data")
	port := flag.Int("port", 8080, "Port to listen to for form submissions")
	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

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

		err := sendWebhook(*url, *nameKey, *messageKey, name, message)
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

	log.Fatal(r.Run(":" + strconv.Itoa(*port)))
}

func sendWebhook(url string, nameKey string, messageKey string, name string, message string) error {
	body := make(map[string]string)
	body[nameKey] = name
	body[messageKey] = message

	jsonBody, err := json.MarshalIndent(body, "", "   ")
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(jsonBody))

	return err
}
