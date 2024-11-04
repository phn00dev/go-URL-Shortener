package httpClient

import (
	"net/http"
	"time"
)

func NewHttp() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}
