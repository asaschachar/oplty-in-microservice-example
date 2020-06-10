# Go Service

## Pre-requisites
- Install [Golang v1.13+](https://golang.org/dl/)
- Clone [this repo](https://github.com/asaschachar/oplty-in-microservice-example)

## Run the Service
1. Open a new terminal window
2. Navigate to the `go-service` directory of this repo:
```bash
cd go-service
```
3. Install the dependencies:
```bash
go get -u github.com/gin-gonic/gin
```

3. Run the service with:
```bash
go run app.go
```

## Evaluate the feature flag
- Replace the `handleRequest` function with the following and update `<Your-SDK-Key>`:
```go
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
```

## Complete Example
- The complete example code for app.go is below:
```go
package main

import "github.com/gin-gonic/gin"
import (
  "net/http"
  "encoding/json"
  "bytes"
)

func handleRequest(c *gin.Context) {
  const sdkKey = "DHbTLoxuXmGPHCTGbrSGKP"

  userObj := map[string]string{"userId": "user123"}
  jsonString, _ := json.Marshal(userObj)

  req, _ := http.NewRequest("POST", "http://localhost:8080/v1/activate", bytes.NewBuffer(jsonString))

  q := req.URL.Query()
  q.Add("featureKey", "hello_world")

  req.URL.RawQuery = q.Encode()
  req.Header.Add("X-Optimizely-SDK-Key", sdkKey)

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

  c.String(200, "Go Service: " + message)
}

func main() {
  r := gin.Default()
  r.GET("/", handleRequest)
  r.Run(":3003")
}
```
