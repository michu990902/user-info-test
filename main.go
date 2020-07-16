package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", indexHandler)

	r.Run(":5000")
}

func indexHandler(c *gin.Context){
	ip := c.ClientIP()
	userInfo, _ := c.Request.Header["User-Agent"]
	

	c.JSON(200, gin.H{
		"ip": ip,
		"user info": userInfo,
	})
}