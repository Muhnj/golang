package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
)

type CallbackRequest struct {
	Method            string `xml:"method"`
	MSISDN            string `xml:"msisdn"`
	TransactionID     string `xml:"transactionId"`
	ReferenceID       string `xml:"referenceId"`
	Amount            string `xml:"amount"`
	ClientTransaction string `xml:"client_transaction"`
	Message           string `xml:"message"`
}

func PayleoCallback(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CallbackRequest

	if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid xml",
			http.StatusBadRequest,
		)
		return
	}

	log.Printf("Callback received: %+v", req)

	response := `<?xml version="1.0" encoding="UTF-8"?>
<request>
<code>200</code>
<message>Transaction received</message>
</request>`

	w.Header().Set(
		"Content-Type",
		"application/xml",
	)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}