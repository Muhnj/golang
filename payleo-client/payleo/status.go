package payleo

import "encoding/json"

func CheckStatus(
	baseURL,
	msisdn,
	merchantCode,
	transactionID,
	consumerKey,
	consumerSecret string,
) (*CheckStatusResponse, error) {

	signURL := baseURL + "/deposit/checkStatus"

	signature := CheckStatusSignature(
		signURL,
		msisdn,
		merchantCode,
		transactionID,
		consumerSecret,
	)

	payload := map[string]any{
		"msisdn":         msisdn,
		"merchantCode":   merchantCode,
		"transactionId":  transactionID,
		"consumerKey":    consumerKey,
		"auth_signature": signature,
	}

	respBytes, err := Post(
		signURL+"/",
		payload,
	)

	if err != nil {
		return nil, err
	}

	var resp CheckStatusResponse

	err = json.Unmarshal(respBytes, &resp)

	return &resp, err
}