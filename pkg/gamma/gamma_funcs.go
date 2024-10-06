package gamma

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	ds "github.com/Evan-Whitehouse/go-order-utils/pkg/datastructures"
)

const (
	BaseUrl = "https://gamma-api.polymarket.com"
	Markets = "markets"
	Events  = "events"
)

func GetAllEvents() []ds.Event {
	// Get all events that can be traded on Polymarket
	url := fmt.Sprintf("%s/%s", BaseUrl, Events)

	// add queries
	queries := map[string]string{
		"closed":          "false",
		"enableOrderBook": "true",
	}

	fullUrl := addQueries(url, queries)

	// Make the request
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var events []ds.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal(err)
	}

	return events
}

func addQueries(baseUrl string, queries map[string]string) string {
	// Add queries to the request
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	for key, value := range queries {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
