package reddit

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetByUrl initiates a simple GET request against the Reddit API using the provided URL.
func GetByUrl(url string) (*http.Response, error) {
	hc := getHttpClient()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", generateUserAgent())
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetByUrlWithRetry wraps the GetByUrl call within retry logic for re-attempting GET calls that have the potential to
// randomly error out but return successful on a subsequent call.
func GetByUrlWithRetry(url string, retries int64) (*http.Response, error) {
	var c int64 = 0
	var err error

	for c < retries {
		r, err := GetByUrl(url)
		if err != nil {
			c++
		} else if r.StatusCode == 404 {
			// If the URL 404s, then avoid retrying unnecessarily.
			return nil, fmt.Errorf("404: url `%s` not found", url)
		} else {
			return r, nil
		}

		// If the call will be attempted again, pause briefly before proceeding.
		if c < retries {
			log.Println("call failed. Pausing then retrying.")
			time.Sleep(1 * time.Second)
			log.Println("pause complete. Resuming.")
		}
	}

	return nil, fmt.Errorf("retry count exceeded on url `%s` with error: %s", url, err)
}

// getHttpClient initializes and returns an HTTP client instance.
func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: true,
		},
	}
}
