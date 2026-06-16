package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() {
	_ = godotenv.Load()
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func buildBasicToken(accountSID, secret string) string {
	raw := accountSID + ":" + secret
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

func validateMSISDN(msisdn string) {
	accountSID := getEnv("ACCOUNT_SID")
	secret := getEnv("SECRET")

	apiURL := getEnv("API_URL")
	if apiURL == "" {
		apiURL = "https://validator.trueafrican.com/api/validation/msisdn/request"
	}

	if accountSID == "" || secret == "" {
		fmt.Println("Missing ACCOUNT_SID or SECRET in environment variables")
		return
	}

	token := buildBasicToken(accountSID, secret)

	payload := map[string]string{
		"msisdn": msisdn,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("\nStatus:", resp.Status)
	fmt.Println("Response:", string(body))
}

func main() {
	loadEnv()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter MSISDN (e.g. 256777123456): ")
	msisdn, _ := reader.ReadString('\n')

	msisdn = strings.TrimSpace(msisdn)

	if msisdn == "" {
		fmt.Println("MSISDN cannot be empty")
		return
	}

	validateMSISDN(msisdn)
}