from flask import Flask
import requests
import json

app = Flask(__name__)

@app.route('/')
def hello_world():
    sdk_key = 'DHbTLoxuXmGPHCTGbrSGKP'

    s = requests.Session()
    s.headers.update({'X-Optimizely-SDK-Key': sdk_key})

    payload = {
      "userId": "user123",
    }

    params = {
      "featureKey": "hello_world"
    }

    resp = s.post(url = 'http://localhost:8080/v1/activate', params=params, json=payload)
    feature_result = resp.json()[0]

    print(json.dumps(feature_result, indent=2))

    feature_text = "Feature is On!" if feature_result['enabled'] else "Feature is off :("

    return 'Hello, World! ' + feature_text

if __name__ == '__main__':
      app.run(host='0.0.0.0', port=3002)
