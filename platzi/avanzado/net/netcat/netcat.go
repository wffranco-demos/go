package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	port = flag.Int("p", 8080, "The port to scan")
	host = flag.String("h", "localhost", "The host to scan")
)

func main() {
	flag.Parse()
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *host, *port), 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to", conn.RemoteAddr())
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
