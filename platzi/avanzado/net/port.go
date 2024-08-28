package main

import (
	"fmt"
	"net"
	"time"
)

func OpenPort(site string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", site, port), 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
