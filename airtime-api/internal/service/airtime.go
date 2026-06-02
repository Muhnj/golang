package service

import (
	"encoding/xml"

	"github.com/Munir/airtime-api/internal/trueafrican"
)

type AirtimeService struct {
	Client *trueafrican.Client
}

func (s *AirtimeService) Purchase(
	msisdn string,
	network string,
	value string,
	transaction string,
) (*trueafrican.PurchaseResponse, error) {

	req := trueafrican.AirtimeRequest{
		Method:      "purchaseAirtime",
		MSISDN:      msisdn,
		Network:     network,
		Value:       value,
		Transaction: transaction,
	}

	respBytes, err := s.Client.Send(req)

	if err != nil {
		return nil, err
	}

	var resp trueafrican.PurchaseResponse

	err = xml.Unmarshal(respBytes, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}