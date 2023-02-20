package reddit

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// generateUserAgent creates a string to be used as the User Agent header value when making requests to Reddit's API.
// When using the Reddit API, a User Agent *must* be set.
func generateUserAgent() string {
	return fmt.Sprintf("reddit-go-client-%s", GetRandomStringOfLength(6))
}

// GetRandomStringOfLength generates and returns a random string of length n.
// Logic is based on the following solution: https://stackoverflow.com/a/31832326
func GetRandomStringOfLength(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
