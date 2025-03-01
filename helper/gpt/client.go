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
	- Infer reasonable categories and priorities.  
	- Set DueDate based on the following rules:
		2) If user input specifies just a time like '8 AM' or '5pm', set DueDate to **today's date** with the given time.  
		1) If user input contains days like 'Monday' or 'Sunday', set DueDate to the nearest future day, or today if it's the same day, with time set to 12:00 PM.  
		3) If user input mentions an old date or 'yesterday', set DueDate to the default: "0001-01-01T00:00:00Z".  
		4) If user input does not mention any date, day, or time, set DueDate to the default: "0001-01-01T00:00:00Z".  
	- Output must be valid JSON only.  
	- Correct grammar issues in descriptions automatically.  
	- Always respond in English regardless of input language.  

	Extract structured tasks from user input as valid JSON (no explanations or code blocks):  
	[{
	"ID": integer,        // Unique ID (starting from 1)  
	"Description": string, // Task details  
	"Category": string,    // Work, Personal, etc.  
	"Priority": string,    // "High", "Medium" (default), "Low"  
	"DueDate": string      // Go time.Time format or "0001-01-01T00:00:00Z" as per rules above  
	}]

	**Example:**  
	User: "Prepare a report by Monday, send it to the manager, then follow up."  
	AI Output:  
	[
	{"ID":1, "Description":"Prepare a report", "Category":"Work", "Priority":"High", "DueDate":"2025-03-03T12:00:00Z"},  
	{"ID":2, "Description":"Send report to manager", "Category":"Work", "Priority":"High", "DueDate":"2025-03-03T12:00:00Z"},  
	{"ID":3, "Description":"Follow up with manager", "Category":"Work", "Priority":"Medium", "DueDate":"0001-01-01T00:00:00Z"}  
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
