# Implementing Feature Flags in a Microservice

## Pre-requisites
- Create a free [Optimizely Account](https://www.optimizely.com/rollouts-signup/?utm_source=youtube&utm_campaign=microservice-feature-flags)
- Install [Python 3](https://www.python.org/downloads/)
- Install [Golang 1.13+](https://golang.org/dl/)

## Run Optimizely as a Service
1. Clone the [Optimizely Agent](https://github.com/optimizely/agent/tree/v1.1.0) repo
2. Change to the agent directory with `cd agent`
3. Setup the service with:
```bash
make setup
```
4. For debugging during development, change the polling interval in the `config.yaml` to `10s`
5. Run the service with:
```bash
make run
``` 

## Run an example Python microservice with [Flask](https://flask.palletsprojects.com/en/1.1.x/)
1. In a separate terminal window, create a Python virtual environment with:
```bash
python3 -m venv venv
```
2. Activate our new virtual environment with:
```bash
. venv/bin/activate
```
3. Install the depencies for our application
```bash
pip install Flask
pip install requests
```
4. Create a file `app.py` and paste in the following code to create our minimal Flask service:
```python
from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'
```

5. Run our service with:
```bash
flask run
```

## Implement a feature flag in our Python microservice
1. Set your SDK Key as a header for each request:
```python
import json
import requests

sdk_key = 'DHbTLoxuXmGPHCTGbrSGKP'
s = requests.Session()
s.headers.update({'X-Optimizely-SDK-Key': sdk_key})
```
2. Create a feature flag named `hello_world` in the Optimizely application
3. Make a request for the feature flag in the Python service:
```python
    payload = {
      "userId": "user123",
    }

    params = {
      "featureKey": "hello_world"
    }

    resp = s.post(url = 'http://localhost:8080/v1/activate', params=params, json=payload)
    feature_result = resp.json()[0]

    print(json.dumps(feature_result, indent=2))
```
4. Implement the flag
```python
    feature_text = "Feature is On!" if feature_result['enabled'] else "Feature is off :("
  
    return 'Hello, World! ' + feature_text
```
5. Turn the flag on and off.
6. Profit 🎉
