package src

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reddit-sync/src/reddit"
)

func GetPost() {
	url := "https://www.reddit.com/r/dbz/comments/znnbh4/universe_7_ready_to_win_the_next_tournament_of/.json"
	resp := reddit.GetByUrl(url)
	defer resp.Body.Close()

	var listings []reddit.Listing
	err := json.NewDecoder(resp.Body).Decode(&listings)
	if err != nil {
		log.Fatalf("unable to parse Listings from JSON Response URL `%s`: %s", url, err)
	}

	reddit.ExtractDataFromListings(listings)
}

func debugRawResponse(resp *http.Response) {
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Unable to read JSON response: %s", err)
		}

		log.Println(string(body))
		log.Fatalf("API call did not return 200 for http call. Received: %d", resp.StatusCode)
	}
}
