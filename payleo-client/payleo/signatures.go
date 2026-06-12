package payleo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func hmacSHA256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func DepositSignature(
	url,
	msisdn,
	amount,
	merchantCode,
	transactionID,
	narration,
	secret string,
) string {

	data :=
		strings.TrimSpace(url) + "&" +
			strings.TrimSpace(msisdn) + "&" +
			strings.TrimSpace(amount) + "&" +
			strings.TrimSpace(merchantCode) + "&" +
			strings.TrimSpace(transactionID) + "&" +
			strings.TrimSpace(narration)

	return hmacSHA256(data, secret)
}

func CheckStatusSignature(
	url,
	msisdn,
	merchantCode,
	transactionID,
	secret string,
) string {

	data :=
		strings.TrimSpace(url) + "&" +
			strings.TrimSpace(msisdn) + "&" +
			strings.TrimSpace(merchantCode) + "&" +
			strings.TrimSpace(transactionID)

	return hmacSHA256(data, secret)
}

func InitiateCallbackSignature(
	url,
	msisdn,
	merchantCode,
	transactionID,
	mnoTransID,
	secret string,
) string {

	data :=
		strings.TrimSpace(url) + "&" +
			strings.TrimSpace(msisdn) + "&" +
			strings.TrimSpace(merchantCode) + "&" +
			strings.TrimSpace(transactionID) + "&" +
			strings.TrimSpace(mnoTransID)

	return hmacSHA256(data, secret)
}