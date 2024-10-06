package main

import (
	"fmt"

	"github.com/Evan-Whitehouse/go-order-utils/pkg/gamma"
)

func main() {
	events := gamma.GetAllEvents()
	fmt.Printf("found %d events\n", len(events))
	for _, event := range events {
		fmt.Printf("Event: %s\n Volume: %f\n", event.Ticker, event.Volume)
		fmt.Printf("Description: %s\n", event.Description)
		fmt.Println("-------------------------------------------------")
	}
	fmt.Println(events[0].Description)
}
