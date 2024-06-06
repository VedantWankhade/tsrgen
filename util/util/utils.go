package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetBearerToken(username, token string) string {
	creds := fmt.Sprintf("%s:%s", username, token)
	return base64.StdEncoding.EncodeToString([]byte(creds))
}

func MakeHTTPRequest(method, url string, headers map[string]string, body io.Reader, params map[string]string) (io.ReadCloser, error) {
	if len(params) != 0 {
		var paramsString bytes.Buffer
		paramsString.WriteRune('?')
		for k, v := range params {
			paramsString.WriteString(fmt.Sprintf("%s=%s", k, v))
		}
		url += paramsString.String()
	}
	log.Printf("Making request: %s %s\n", method, url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request object failed: %w", err)
	}
	client := &http.Client{}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return nil, fmt.Errorf("http request failed: status: %v(%v): err: %w", res.StatusCode, res.Status, err)
	}
	return res.Body, nil
}
