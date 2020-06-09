const express = require('express')
const querystring = require('querystring');
const httpRequest = require('axios')
const app = express()
const port = 3001

async function handleRequest(req, res) {
  res.send('JavaScript Service: ' + 'Hello World!')
}

app.get('/', handleRequest)

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
