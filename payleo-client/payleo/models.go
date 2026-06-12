package payleo

type DepositRequest struct {
	MSISDN        string `json:"msisdn"`
	Amount        string `json:"amount"`
	MerchantCode  string `json:"merchantCode"`
	TransactionID string `json:"transactionId"`
	ConsumerKey   string `json:"consumerKey"`
	AuthSignature string `json:"auth_signature"`
	Narration     string `json:"narration"`
}

type APIResponse struct {
	Status        string `json:"status"`
	Code          int    `json:"code"`
	Message       string `json:"message"`
	TransactionID string `json:"transactionId"`
}

type CheckStatusResponse struct {
	Status              string `json:"status"`
	Code                int    `json:"code"`
	TransactionID       string `json:"transactionId"`
	ClientTransactionID string `json:"clientTransactionId"`
	MNOTransID          any    `json:"mnoTransId"`
}