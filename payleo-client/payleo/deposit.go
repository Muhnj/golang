package payleo

import (
	"encoding/json"
)

func Deposit(
	baseURL,
	msisdn,
	amount,
	merchantCode,
	consumerKey,
	consumerSecret,
	transactionID,
	narration string,
) (*APIResponse, error) {

	signURL := baseURL + "/deposit"

	signature := DepositSignature(
		signURL,
		msisdn,
		amount,
		merchantCode,
		transactionID,
		narration,
		consumerSecret,
	)

	payload := DepositRequest{
		MSISDN:        msisdn,
		Amount:        amount,
		MerchantCode:  merchantCode,
		TransactionID: transactionID,
		ConsumerKey:   consumerKey,
		AuthSignature: signature,
		Narration:     narration,
	}

	respBytes, err := Post(
		signURL+"/",
		payload,
	)

	if err != nil {
		return nil, err
	}

	var resp APIResponse

	err = json.Unmarshal(respBytes, &resp)

	return &resp, err
}