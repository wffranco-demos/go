package main

import (
	"flag"
	"fmt"
	"sync"
	"testing"
)

/**
 * Sites:
 * - scanme.nmap.org
 * - scanme.webscantest.com
 */
var site = flag.String("site", "scanme.nmap.org", "The site to scan")

func TestAvailablePorts(t *testing.T) {
	flag.Parse()
	fmt.Println("Site:", *site)
	var wg sync.WaitGroup

	for p := 0; p < 65535; p++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			if OpenPort(*site, port) {
				fmt.Printf("Port %d is open\n", port)
			}
		}(p)
	}

	wg.Wait()
}
