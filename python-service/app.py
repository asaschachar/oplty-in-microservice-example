
from flask import Flask
import requests
import json

app = Flask(__name__)

@app.route('/')
def handle_request():
    return 'Python Service: ' + 'Hello World!'

if __name__ == '__main__':
      app.run(host='0.0.0.0', port=3002)
