# JavaScript Service

## Pre-requisites
- Install [Node v10+](https://nodejs.org/en/download/)
- Clone [this repo](https://github.com/asaschachar/oplty-in-microservice-example)

## Run the Service
1. Open a new terminal window
2. Navigate to the `javascript-service` directory of this repo:
```bash
cd javascript-service
```

3. Install dependencies with:
```bash
npm install
```

4. Run the service with:
```bash
node app.js
```

## Evaluate the feature flag
- Replace the `handleRequest` function with the following and update `<Your-SDK-Key>`:

```javascript
async function handleRequest(req, res) {
  response = await httpRequest({
    url: 'http://localhost:8080/v1/activate',
    method: 'POST',
    headers: { 'X-Optimizely-SDK-Key': '<Your-SDK-Key>' },
    params: { featureKey: 'hello_world' },
    data: JSON.stringify({ userId: 'user123'}),
  });
 
  let enabled = Boolean(response.data[0].enabled)
  let message = enabled
    ? "Feature is ON!"
    : "Feature is off :("

  res.send('JavaScript Service: ' + message)
}
```

## Complete Example
- The complete example code for app.js is below:
```javascript
const express = require('express')
const querystring = require('querystring');
const axios = require('axios')
const app = express()
const port = 3001

async function handleRequest(req, res) {
  response = await axios({
    method: 'POST',
    url: 'http://localhost:8080/v1/activate',
    headers: { 'X-Optimizely-SDK-Key': 'DHbTLoxuXmGPHCTGbrSGKP' },
    params: { featureKey: 'hello_world' },
    data: JSON.stringify({ userId: 'user123'}),
  });
 
  let enabled = Boolean(response.data[0].enabled)
  let message = enabled
    ? "Feature is ON!"
    : "Feature is off :("

  res.send('JavaScript Service: ' + message)
}

app.get('/', handleRequest)

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
```
