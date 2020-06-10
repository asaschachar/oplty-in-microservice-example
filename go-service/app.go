package main

import "github.com/gin-gonic/gin"

import "github.com/imroc/req"
import "github.com/optimizely/agent/pkg/optimizely"
import (
  "encoding/json"
)

func handleRequest(c *gin.Context) {
  c.String(200, "Go Service: " + "Hello World!")
}

func main() {
  r := gin.Default()
  r.GET("/", handleRequest)
  r.Run(":3003")
}
