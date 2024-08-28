package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	chat            = make(chan string)

	host = flag.String("h", "localhost", "The host to scan")
	port = flag.Int("p", 8080, "The port to scan")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	client := conn.RemoteAddr().String()
	fmt.Printf("Connected to client: %s\n", client)
	message := make(chan string)
	go MessageWrite(conn, message)

	message <- fmt.Sprintf("Welcome %s", client)
	chat <- fmt.Sprintf("New user has arrived: %s", client)
	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		chat <- fmt.Sprintf("%s: %s\n", client, inputMessage.Text())
	}

	leavingClients <- message
	chat <- fmt.Sprintf("User has left: %s", client)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-chat:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-incomingClients:
			clients[cli] = true
		case cli := <-leavingClients:
			delete(clients, cli)
			close(cli)
		}
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chat server is running on", listener.Addr())
	go Broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go HandleConnection(conn)
	}
}
