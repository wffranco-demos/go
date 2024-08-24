package net

import (
	"fmt"
	"testing"
)

func TestAvailablePorts(t *testing.T) {

	for port := 0; port < 100; port++ {
		if OpenPort(port) {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}
