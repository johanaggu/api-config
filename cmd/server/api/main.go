package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	api := gin.Default()
	api.GET("/ping", func(c *gin.Context) { 
		c.String(http.StatusOK, "pong")
	})
	api.Run()
}