package main

import (
	"github.com/gin-gonic/gin"
)

func uploadInitHandler(c *gin.Context) {
	c.Writer.WriteString("Hello")
}

func main() {
	r := gin.Default()
	r.GET("/v1/upload/init", uploadInitHandler)
	r.Run(":8080")
}