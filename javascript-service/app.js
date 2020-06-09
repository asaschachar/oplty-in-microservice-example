const express = require('express')
const querystring = require('querystring');
const axios = require('axios')
const app = express()
const port = 3001

app.get('/', async (req, res) => {
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
})

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))
