package payleo

func InitiateCallback(
	baseURL,
	msisdn,
	merchantCode,
	transactionID,
	mnoTransID,
	consumerKey,
	consumerSecret string,
) ([]byte, error) {

	signURL := baseURL + "/deposit/initiateCallback"

	signature := InitiateCallbackSignature(
		signURL,
		msisdn,
		merchantCode,
		transactionID,
		mnoTransID,
		consumerSecret,
	)

	payload := map[string]any{
		"msisdn":         msisdn,
		"merchantCode":   merchantCode,
		"transactionId":  transactionID,
		"consumerKey":    consumerKey,
		"auth_signature": signature,
		"mnoTransId":     mnoTransID,
	}

	return Post(signURL+"/", payload)
}