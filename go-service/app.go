package main

import "github.com/gin-gonic/gin"
import (
  "net/http"
  "encoding/json"
  "bytes"
)

func handleRequest(c *gin.Context) {

  userObj := map[string]string{"userId": "user123"}
  jsonString, _ := json.Marshal(userObj)

  req, _ := http.NewRequest("POST", "http://localhost:8080/v1/activate", bytes.NewBuffer(jsonString))

  q := req.URL.Query()
  q.Add("featureKey", "hello_world")

  req.URL.RawQuery = q.Encode()
  req.Header.Add("X-Optimizely-SDK-Key", "DHbTLoxuXmGPHCTGbrSGKP")

  client := &http.Client{}
  resp, _ := client.Do(req)

  defer resp.Body.Close()

  var results []map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&results)

  var enabled bool
  enabled = results[0]["enabled"].(bool)

  var message string
  if enabled {
    message = "Feature is ON!"
  } else {
    message = "Feature is off :("
  }

  c.JSON(200, gin.H{
    "message": message,
  })
}

func main() {
  r := gin.Default()
  r.GET("/", handleRequest)
  r.Run(":3003")
}
