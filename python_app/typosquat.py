#!/usr/bin/python3

import requests
import sys

url = "https://api.github.com/events"
response = requests.get(url)

# Accessing response data
print(f"Status Code: {response.status_code}")

sys.exit(0)
