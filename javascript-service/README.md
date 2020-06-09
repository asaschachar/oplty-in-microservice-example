# JavaScript Service

## Pre-requisites
- Install [Node v10+](https://nodejs.org/en/download/)

## Run the Service
1. Open a new terminal window
2. Navigate to the `javascript-service` directory:
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
  const SDK_KEY = '<Your-SDK-Key>'

  let params = {
    featureKey: 'hello_world'
  }

  let payload = {
    userId: 'user123'
  }

  let response = await axios({
    method: 'POST',
    headers: { 'X-Optimizely-SDK-Key': SDK_KEY },
    url: 'http://localhost:8080/v1/activate',
    params: params,
    data: JSON.stringify(payload),
  });

  console.log(response.data);

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
  const SDK_KEY = 'DHbTLoxuXmGPHCTGbrSGKP'

  let params = {
    featureKey: 'hello_world'
  }

  let payload = {
    userId: 'user123'
  }

  let response = await axios({
    method: 'POST',
    headers: { 'X-Optimizely-SDK-Key': SDK_KEY },
    url: 'http://localhost:8080/v1/activate',
    params: params,
    data: JSON.stringify(payload),
  });

  console.log(response.data);

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
