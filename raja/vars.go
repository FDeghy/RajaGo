package raja

import (
	"crypto/tls"
	"net/http"
	"time"
)

const (
	BASE_URL    = "https://www.raja.ir"
	SERVICE_URL = "https://hostservice.raja.ir"
	USER_AGENT  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
)

var (
	Client = http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
)
