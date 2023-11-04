import http.client
import json
from dotenv import load_dotenv
import os

# Load environment variables from .env file
load_dotenv()

# Access the BITAPAI_API_KEY from the environment variables
bitapai_key = os.getenv('BITAPAI_API_KEY')

print(bitapai_key)

conn = http.client.HTTPSConnection("api.bitapai.io")
payload = json.dumps({
  "messages": [
    {
      "role": "user",
      "content": "What is the meaning of life?"
    },
  ],
  "pool_id": 4,
  "count": 5,
  "return_all": True
})
headers = {
  'Content-Type': 'application/json',
  'X-API-KEY': str(bitapai_key),  # Use the environment variable here
}
conn.request("POST", "/text", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))
