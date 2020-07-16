package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func main() {
	r := gin.Default()

	r.GET("/", indexHandler)

	r.Run(":5000")
}

func indexHandler(c *gin.Context){
	ip := c.ClientIP()
	userInfo, _ := c.Request.Header["User-Agent"]
	
	ua := user_agent.New(userInfo[0])
	browser, version := ua.Browser()
	
	c.JSON(200, gin.H{
		"ip": ip,
		"os": ua.OS(),
		"browser": browser + " " + version,
	})
}