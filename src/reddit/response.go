package reddit

import (
	"encoding/json"
	"log"
)

type PostData struct {
	Title string `json:"title"`
}

type ListingChild struct {
	Kind string           `json:"kind"`
	Data *json.RawMessage `json:"data"`
}

type ListingData struct {
	After    string         `json:"after"`
	Dist     int64          `json:"dist"`
	Children []ListingChild `json:"children"`
}

type Listing struct {
	Kind string           `json:"kind"`
	Data *json.RawMessage `json:"data"`
}

// ExtractDataFromListings verifies the provided slice of Listings match the expected Kind and then forwards them to
// the Child data extractor.
func ExtractDataFromListings(listings []Listing) {
	for _, listing := range listings {
		if listing.Kind == "Listing" {
			var listingData ListingData
			if err := json.Unmarshal(*listing.Data, &listingData); err != nil {
				log.Fatalf("unable to parse Listing data: %s", err)
			}

			for _, listingChild := range listingData.Children {
				extractDataFromListingChild(listingChild)
			}
		} else {
			log.Fatalf("unexpected Listing kind: %s", listing.Kind)
		}
	}
}

// extractDataFromListingChild determines the Child's Kind and extract its data accordingly.
func extractDataFromListingChild(child ListingChild) {
	if child.Kind == "t3" {
		var postData PostData
		if err := json.Unmarshal(*child.Data, &postData); err != nil {
			log.Fatalf("unable to parse Post data from Child object: %s", err)
		}

		processPost(postData)
	} else if child.Kind == "t1" {
		log.Println("Process Comment")
	} else if child.Kind == "more" {
		log.Println("Process more")
	} else {
		log.Fatalf("Unexpected Type: %v", child)
	}
}

// processPost analyzes and processes the provided Post data.
func processPost(postData PostData) {
	log.Println("Actual ListingChild Data: ", postData.Title)
}
