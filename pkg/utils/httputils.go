package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, headers map[string]string, body interface{}, response interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return ParseResponse(res, response)
}

func Get(url string, headers map[string]string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return ParseResponse(res, response)
}

func DeleteRequest(url string, headers map[string]string, response interface{}) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return ParseResponse(res, response)
}

func ParseResponse(res *http.Response, response interface{}) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if response != nil {
			return json.Unmarshal(body, response)
		}
		return nil
	}
	return fmt.Errorf("HTTP error: %s", res.Status)
}
