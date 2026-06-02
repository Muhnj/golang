package trueafrican

import "encoding/xml"

type PurchaseResponse struct {
	XMLName    xml.Name `xml:"Response"`
	ResponseID string   `xml:"responseId"`
	Status     string   `xml:"status"`
	Code       string   `xml:"code"`
}

type BalanceResponse struct {
	XMLName xml.Name `xml:"Response"`
	Code    string   `xml:"code"`
	Status  string   `xml:"status"`
	Balance string   `xml:"balance"`
}