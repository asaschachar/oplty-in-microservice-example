# Python Service

## Pre-requisites
- Install [Python 3](https://www.python.org/downloads/)
- Clone [this repo](https://github.com/asaschachar/oplty-in-microservice-example)

## Run the Service
1. Open a new terminal window
2. Navigate to the `python-service` directory of this repo:
```bash
cd python-service
```

3. Create a virtual environment:
```bash
python3 -m venv venv
```

4. Activate the virtual environment:
```bash
. venv/bin/activate
```

5. Install the dependencies:
```bash
pip install -r requirements.txt
```

6. Run the service with:
```bash
python app.py
```

## Evaluate the feature flag
- Replace the `handle_request` function with the following and update `<Your-SDK-Key>`:
```python
def handle_request():
    s = requests.Session()
    resp = s.post(
      url='http://localhost:8080/v1/activate',
      headers={'X-Optimizely-SDK-Key': 'XBAY4tRMRJwYUmyuqcprpP'},
      params={ 'featureKey': 'hello_world'},
      json={ 'userId': 'user123' }
    )

    feature_result = resp.json()[0]
    feature_text = 'Feature is On!' if feature_result['enabled'] else 'Feature is off :('
    return 'Python Service: ' + feature_text
```

## Complete Example
- The complete example code for app.py below:
```python
from flask import Flask
import requests
import json

app = Flask(__name__)

@app.route('/')
def handle_request():
    s = requests.Session()
    resp = s.post(
      url='http://localhost:8080/v1/activate',
      headers={'X-Optimizely-SDK-Key': 'XBAY4tRMRJwYUmyuqcprpP'},
      params={ 'featureKey': 'hello_world'},
      json={ 'userId': 'user123' }
    )

    feature_result = resp.json()[0]
    feature_text = 'Feature is On!' if feature_result['enabled'] else 'Feature is off :('
    return 'Python Service: ' + feature_text 

if __name__ == '__main__':
      app.run(host='0.0.0.0', port=3002)
```
