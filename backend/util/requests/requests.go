package requests

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/bytedance/sonic"
)

// Type alias for a map of strings to make it easier to define headers
type Headers = map[string]string

func PostRequestURL(url string, body interface{}) (Map, error) {
	return PostRequestWithHeaders[Map](url, body, Headers{})
}

func PostRequestWithHeaders[T any](url string, body interface{}, headers Headers) (T, error) {
	// Declared here so it can be returned as nil before it's actually used
	var data T

	// Encode body to JSON
	byteBody, err := sonic.Marshal(body)
	if err != nil {
		return data, err
	}

	// Set headers
	reqHeaders := http.Header{}
	reqHeaders.Set("Content-Type", "application/json")
	for key, value := range headers {
		reqHeaders.Set(key, value)
	}

	// Send the request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byteBody))
	if err != nil {
		log.Fatalln("failed to create request:", err)
	}
	req.Header = reqHeaders
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	// Grab all bytes from the buffer
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return data, err
	}

	// Parse body into JSON
	err = sonic.Unmarshal(buf.Bytes(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
