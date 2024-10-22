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
		fmt.Printf("MARKETS\n")
		for _, market := range event.Markets {
			fmt.Printf("Q: %s\n", market.Question)
			fmt.Printf("Market: %v\n", market.Outcomes)
			fmt.Printf("Volume: %f\n", market.Volume)
			fmt.Printf("Clob token IDs:\n%v\n", market.ClobTokenIds)
			break
		}
		fmt.Println("-------------------------------------------------")
		break
	}
	fmt.Println(events[0].Description)
}
