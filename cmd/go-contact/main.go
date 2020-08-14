package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Println(name, message)
	})

	_ = r.Run()
}
