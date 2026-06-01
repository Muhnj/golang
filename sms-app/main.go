package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"bytes"
	"encoding/json"

	"github.com/joho/godotenv"
)

type SMSRequest struct {
	Msisdn  []string `json:"msisdn"`
	Message string   `json:"message"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}
func main() {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load .env file")
		return
	}

	// SMS credentials
	username := os.Getenv("SMS_USERNAME")
	password := os.Getenv("SMS_PASSWORD")
	apiURL := os.Getenv("SMS_API_URL")

	var clientName string
	var msisdn string

	reader := bufio.NewReader(os.Stdin)

	//client name
	fmt.Print("Enter client name: ")
	fmt.Scanln(&clientName)

	//Phone number
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&msisdn)

	// SMS message
	fmt.Print("Enter message: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message")
		return
	}

	message = strings.TrimSpace(message)

	// summary message 
	fmt.Println("\n----- Welcome To True Africa Limited -----")
	fmt.Println("Client:", clientName)
	fmt.Println("Phone:", msisdn)
	fmt.Println("Message:", message)

	// confirmation 
	var choice string

	fmt.Print("\nSend SMS? (y/n): ")
	fmt.Scanln(&choice)

	if strings.ToLower(choice) != "y" {
		fmt.Println("SMS cancelled.")
		return
	}

	// Call SMS function
	sendSMS(
		apiURL,
		username,
		password,
		msisdn,
		message,
	)
}

func sendSMS(apiURL, username, password, msisdn, message string) {

	fmt.Println("\nSending SMS...")

	// Build request body
	payload := SMSRequest{
		Msisdn:   []string{msisdn},
		Message:  message,
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	// Set headers (VERY IMPORTANT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body := make([]byte, 1024)
	resp.Body.Read(body)

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response:", string(body))
}