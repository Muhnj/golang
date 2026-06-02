package trueafrican

import "encoding/xml"

type AirtimeRequest struct {
	XMLName     xml.Name `xml:"request"`
	Username    string   `xml:"username"`
	Password    string   `xml:"password"`
	Method      string   `xml:"method"`
	MSISDN      string   `xml:"msisdn,omitempty"`
	MSISDN2     string   `xml:"msisdn2,omitempty"`
	Network     string   `xml:"network,omitempty"`
	Value       string   `xml:"value,omitempty"`
	Transaction string   `xml:"transaction,omitempty"`
	RequestID   string   `xml:"requestId,omitempty"`
}