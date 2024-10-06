package gamma_test

import (
	"fmt"
	"testing"

	"github.com/Evan-Whitehouse/go-order-utils/gamma"
)

// test get all events
func TestGetAllEvents(t *testing.T) {
	fmt.Print(gamma.GetAllEvents())
}
