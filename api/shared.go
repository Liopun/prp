package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var (
	httpClient = &http.Client{Timeout: 5 * time.Second}
)

type genericResponse struct {
	statusCode int
	body []byte
	error error
}

func makeRequest(reqtype string, path string, body []byte) *genericResponse {
	req, err := http.NewRequest(reqtype, fmt.Sprintf(path, viper.GetString("BASE_URL")), bytes.NewBuffer(body))
	if err != nil {
		return &genericResponse{req.Response.StatusCode, nil, err}
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Cache-Control", "no-cache, no-store")
	req.Header.Set("Authorization", "token "+viper.GetString("token"))

	resp, err := httpClient.Do(req)
	if err != nil {
		return &genericResponse{resp.StatusCode, nil, err}
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &genericResponse{resp.StatusCode, nil, err}
	}


	if resp.StatusCode == 200 || resp.StatusCode == 202 {
		return &genericResponse{resp.StatusCode, respBody, err}
	}

	return &genericResponse{
		resp.StatusCode,
		[]byte{},
		fmt.Errorf("make request error: %v", string(respBody)),
	}

}
