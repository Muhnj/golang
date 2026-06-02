package trueafrican

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type Client struct {
	URL      string
	Username string
	Password string
}

func (c *Client) Send(req AirtimeRequest) ([]byte, error) {

	req.Username = c.Username
	req.Password = c.Password

	xmlData, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		c.URL,
		bytes.NewBuffer(xmlData),
	)

	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "text/xml")
	httpReq.Header.Set("Content-transfer-encoding", "text")

	client := &http.Client{}

	resp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}