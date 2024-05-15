package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func helloWorld() {
	fmt.Println("Hello, World!")
}

func main() {
	r := gin.Default()

	r.GET("/hello-world", func(c *gin.Context) {
		helloWorld()
		c.String(200, "Hello, World!")
	})

	r.Run()
}
