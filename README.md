# Demo of Optimizely in a Microservice Environment

## Step 1: Setup the JavaScript Service

### Pre-requisites
- Install [Node v10+](https://nodejs.org/en/download/)

### Run the Service
1. Open a new terminal window
2. Clone this repository.
3. Navigate to the JavaScript service directory of the repository:
```bash
cd optly-in-microservice-example/javascript-service
```
4. Install dependencies with:
```bash
npm install
```
5. Run the service with:
```bash
node app.js
```
6. Open a browser and navigate to `http://localhost:3001`

## Step 2: Setup the Python Service (optional)

### Pre-requisites
- Install [Python 3](https://www.python.org/downloads/)

### Run the Service
1. Open a new terminal window
2. Clone this repository if you haven't already
3. Navigate to the Python service directory of the repository:
```
cd optly-in-microservice-example/python-service
```
4. Create a virtual environment:
```bash
python3 -m venv venv
```
5. Activate the virtual environment:
```bash
. venv/bin/activate
```
6. Install the dependencies:
```bash
pip install -r requirements.txt
```
7. Run the service with:
```bash
python app.py
```
8. Open a browser and navigate to `http://localhost:3002`

## Step 3: Create a feature flag
1. Create a free [Optimizely Rollouts Account](https://www.optimizely.com/rollouts-signup/?utm_source=youtube&utm_campaign=asa-microservice-feature-flags)
2. Create a feature named `hello_world` in the [Optimizely Application](https://app.optimizely.com)

## Step 4: Setup & Run [Optimizely Agent](https://docs.developers.optimizely.com/full-stack/docs/setup-optimizely-agent)
### Pre-requisites
 - Install [Golang v1.13+](https://golang.org/dl/)

### Run Optimizely as a Service
1. Clone the [Optimizely Agent repo](https://github.com/optimizely/agent/tree/v1.2.0)
2. Change to the agent directory:
 ```bash
 cd agent
 ```
3. [Optional] For faster debugging during development, change the `pollingInterval` in the `config.yaml` to `10s`
4. Setup Optimizely Agent with:
```bash
make setup
```
5. Run Optimizely Agent with:
 ```bash
 make run
 ```
8. Open a browser and navigate to `http://localhost:8080`

## Step 5: Update the Services to Evaluate the Feature Flag
Update one or more of the microservices to request feature flag state from the Optimizley Agent [activate endpoint](https://docs.developers.optimizely.com/full-stack/docs/use-optimizely-agent#section-manage-features): `POST /v1/activate?featureKey={featureKey}`

### For the JavaScript Service:
- Replace the `handleRequest` function in `app.js` with the following and update `<Your-SDK-Key>`:

```javascript
async function handleRequest(req, res) {
  let response = await axios({
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
2. Save your changes and restart the service with Ctrl+C and `node app.js`
3. Refresh your browser pointing to `http://localhost:3001`

### For the Python Service:
1. Replace the `handle_request` function in `app.py` with the following and update `<Your-SDK-Key>`:
```python
def handle_request():
    s = requests.Session()
    resp = s.post(
      url='http://localhost:8080/v1/activate',
      headers={'X-Optimizely-SDK-Key': '<Your-SDK-Key>'},
      params={ 'featureKey': 'hello_world'},
      json={ 'userId': 'user123' }
    )

    feature_result = resp.json()[0]
    feature_text = 'Feature is On!' if feature_result['enabled'] else 'Feature is off :('
    return 'Python Service: ' + feature_text
```
2. Save your changes and restart the service with Ctrl+C and `python app.py`
3. Refresh your browser pointing to `http://localhost:3002`


## Step 6: Turn the feature on!
1. Turn the feature flag on for the environment corresponding to the SDK Key you used above.
2. Look for the console log in Optimizely Agent for when the update occurs
3. Refresh the browser pointing to the services
4. Profit 🎉 ! You've now implemented feature flags in a microservice environment with Optimizley Agent!
