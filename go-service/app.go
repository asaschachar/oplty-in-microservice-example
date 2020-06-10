package main

import "github.com/gin-gonic/gin"
import "github.com/imroc/req"
import "github.com/optimizely/agent/pkg/optimizely"
import (
  "encoding/json"
)

func handleRequest(c *gin.Context) {
  userObj := map[string]string{"userId": "user123"}
  jsonString, _ := json.Marshal(userObj)

  r, _ := req.Post(
    "http://localhost:8080/v1/activate",
    req.Header{"X-Optimizely-SDK-Key": "DHbTLoxuXmGPHCTGbrSGKP"},
    req.BodyJSON(jsonString),
    req.QueryParam{"featureKey": "hello_world"},
  )

  var results []optimizely.Decision
  r.ToJSON(&results)

  var enabled bool = results[0].Enabled

  var message string
  if enabled {
    message = "Feature is ON!"
  } else {
    message = "Feature is off :("
  }

  c.String(200, "Go Service: " + message)
}

func main() {
  r := gin.Default()
  r.GET("/", handleRequest)
  r.Run(":3003")
}
