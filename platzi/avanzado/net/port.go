package net

import (
	"fmt"
	"net"
)

var site = "scanme.nmap.org"

func OpenPort(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", site, port))
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
