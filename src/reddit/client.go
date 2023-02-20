package reddit

import (
	"fmt"
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
		} else {
			return r, nil
		}

		// If the call will be attempted again, pause briefly before proceeding.
		if c < retries {
			time.Sleep(1 * time.Second)
		}
	}

	return nil, fmt.Errorf("retry count exceeded on url `%s` with error: %s", url, err)
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
