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

	// Define the handler function
	http.HandleFunc("/api/query", handleQuery)

	// Start the HTTP server
	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Define a struct for your expected request format
type QueryRequest struct {
	Content string `json:"content"`
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // In production, replace '*' with your frontend's actual URL

	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var queryReq QueryRequest
	err := json.NewDecoder(r.Body).Decode(&queryReq)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Access the BITAPAI_API_KEY from the environment variables
	bitapaiKey := os.Getenv("BITAPAI_API_KEY")

	// Dynamically create the payload using the content from the request
	payload := map[string]interface{}{
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": queryReq.Content,
			},
		},
		"pool_id":     4,
		"count":       5,
		"return_all":  true,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	conn, err := http.NewRequest("POST", "https://api.bitapai.io/text", bytes.NewBuffer(payloadBytes))
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

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode and write the JSON response
	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
