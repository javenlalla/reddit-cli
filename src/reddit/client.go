package reddit

import (
	"log"
	"net/http"
	"time"
)

// GetByUrl initiates a simple GET request against the Reddit API using the provided URL.
func GetByUrl(url string) *http.Response {
	hc := getHttpClient()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", generateUserAgent())
	resp, err := hc.Do(req)
	if err != nil {
		log.Fatalf("get call to %s failed: %s", url, err)
	}

	return resp
}

// getHttpClient initializes and returns an HTTP client instance.
func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 15,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
}
