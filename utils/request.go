package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	Status  string
	Headers map[string]string
	Body    []byte
}

// HttpRequest http request method
func HttpRequest(method string, urlHost string, headers map[string]string, parameters interface{}) (*Response, error) {

	// create json body
	jsonStr, err := EncodeJSON(parameters)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest(method, urlHost, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	// Set request header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Get response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed with status: " + resp.Status + "\n response body: " + string(body))
	}

	return &Response{
		Status:  resp.Status,
		Headers: extractHeaders(resp.Header),
		Body:    body,
	}, nil
}

// HTTPGet http get method
func HTTPGet(urlHost string, headers map[string]string, parameters interface{}) (*Response, error) {
	return HttpRequest("GET", urlHost, headers, parameters)
}

// HTTPPost http get method
func HTTPPost(urlHost string, headers map[string]string, parameters interface{}) (*Response, error) {
	return HttpRequest("POST", urlHost, headers, parameters)
}

// HTTPDelete http get method
func HTTPDelete(urlHost string, headers map[string]string, parameters interface{}) (*Response, error) {
	return HttpRequest("DELETE", urlHost, headers, parameters)
}

// extractHeaders http.Header formatting from map[string][]string to map[string]string
func extractHeaders(h http.Header) map[string]string {
	headers := make(map[string]string)
	for k, v := range h {
		headers[k] = v[0]
	}
	return headers
}
