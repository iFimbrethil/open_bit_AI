package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Access the BITAPAI_API_KEY from the environment variables
	bitapaiKey := os.Getenv("BITAPAI_API_KEY")

	fmt.Println(bitapaiKey)

	payload := []byte(`{
		"messages": [
			{
				"role": "user",
				"content": "What is the meaning of life?"
			}
		],
		"pool_id": 4,
		"count": 5,
		"return_all": true
	}`)

	conn, err := http.NewRequest("POST", "https://api.bitapai.io/text", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	conn.Header.Add("Content-Type", "application/json")
	conn.Header.Add("X-API-KEY", bitapaiKey)

	client := &http.Client{}
	res, err := client.Do(conn)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer res.Body.Close()

	var result map[string]interface{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println(result)
}
