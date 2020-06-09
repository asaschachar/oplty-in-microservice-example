package main

import "github.com/gin-gonic/gin"

func handleRequest(c *gin.Context) {
  c.JSON(200, gin.H{
    "Go Service: ": "Hello World!",
  })
}

func main() {
  r := gin.Default()
  r.GET("/", handleRequest)
  r.Run(":3003")
}
