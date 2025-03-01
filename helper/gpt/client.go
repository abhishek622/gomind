package gpt

import (
	"abhishek622/gomind/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestPayload struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ResponsePayload struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GenerateTask(userInput string) (string, error) {
	apiURL := "https://api.awanllm.com/v1/chat/completions"
	apiKey := utils.AppConfig.AWANLLM_API_KEY

	if apiKey == "" {
		return "", fmt.Errorf("AWANLLM_API_KEY is not set in environment variables")
	}

	prompt := `You are a task management assistant. 

**Rules:**  
- Infer reasonable categories and priorities
- Set DueDate based on user input:
  1. If input contains a day like "Monday" or "tomorrow," set DueDate to the nearest future date. If it's today, use today's date. The time should match what the user specified — if no time is given, default to 12:00 PM.
  2. If a specific time is mentioned (e.g., "7 AM" or "5 PM"), set DueDate to today's date with the given time.
  3. If an old date or "yesterday" is mentioned, use the default DueDate: "0001-01-01T00:00:00Z".
  4. If there's no mention of a day, date, or time, use the default DueDate: "0001-01-01T00:00:00Z".
- Ensure time is correctly parsed and matched to the specified day.
- Output must be valid JSON only
- Automatically correct grammar issues in descriptions
- Always respond in English regardless of input language

Extract structured tasks from user input as valid JSON (no explanations or code blocks):
[{
  "ID": integer,        // Unique ID (starting from 1)
  "Description": string, // Task details
  "Category": string,    // Work, Personal, etc.
  "Priority": string,    // "High", "Medium" (default), "Low"
  "DueDate": string      // Use Go's time.Time format. Example: "2025-03-04T07:00:00Z"
}]

**Example:**  
User: "create the task for tomorrow, Workout at 7 AM, learn Go from 10 AM, and continue the Gomind project at 6 PM"  
AI Output:  
[
  {"ID": 1, "Description": "Workout", "Category": "Health", "Priority": "High", "DueDate": "2025-03-02T07:00:00Z"},
  {"ID": 2, "Description": "Learn Go", "Category": "Education", "Priority": "Medium", "DueDate": "2025-03-02T10:00:00Z"},
  {"ID": 3, "Description": "Continue the Gomind project", "Category": "Work", "Priority": "High", "DueDate": "2025-03-02T18:00:00Z"}
]
`

	// Initialize the request payload
	requestData := RequestPayload{
		Model: "Meta-Llama-3.1-70B-Instruct",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: userInput,
			},
		},
	}

	// Marshal the request data to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check for a successful response
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("error: received status code %d, response: %s", response.StatusCode, body)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response data
	var result ResponsePayload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// Check if there are any choices in the response
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

	return result.Choices[0].Message.Content, nil
}
