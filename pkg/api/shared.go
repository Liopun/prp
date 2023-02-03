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

type requestResponse struct {
	statusCode int
	message string
	error error
}

func makeRequest(reqtype string, path string, body []byte) *requestResponse {
	req, err := http.NewRequest(reqtype, fmt.Sprintf(path, viper.GetString("BASE_URL")), bytes.NewBuffer(body))
	if err != nil {
		return &requestResponse{
			statusCode: req.Response.StatusCode,
			message: "",
			error: err,
		}
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Cache-Control", "no-cache, no-store")
	req.Header.Set("Authorization", "token "+viper.GetString("token"))

	resp, err := httpClient.Do(req)
	if err != nil {
		return &requestResponse{
			statusCode: req.Response.StatusCode,
			message: "",
			error: err,
		}
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &requestResponse{
			statusCode: resp.StatusCode,
			message: "",
			error: err,
		}
	}

	return &requestResponse{
		statusCode: resp.StatusCode,
		message: string(respBody),
		error: nil,
	}
}
