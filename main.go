package main

import (
	"./tests"
	"bytes"
	"fmt"
	"net"
)

func init() {
	fmt.Println("")
	fmt.Println("          ==========")
	fmt.Println("         |  IRCPTS  |")
	fmt.Println("          ==========\n")

	fmt.Println("The IRC protocol used can be found at:")
	fmt.Println("https://tools.ietf.org/html/rfc2812\n")
}

func main() {
	// Connected clients.
	clients := make([]Client, 50)

	// Channel that collects input from all clients
	in := make(chan string)

	// Get input from a client and send it to all the other clients.
	go func() {
		for {
			input := <-in

			for _, client := range clients {
				client.in <- input
			}
		}
	}()

	listen(":6667", &clients, in)
}

// Listen for incoming client connections
func listen(port string, clients *[]Client, in chan string) {
	netlisten, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	defer netlisten.Close()

	fmt.Println("Listening on 127.0.0.1:6667")

	for {
		conn, err := netlisten.Accept()
		if err != nil {
			fmt.Println(err)
		}

		newClient(conn, in, *clients)
	}
}

// A new client as joined the server.
// Run all tests involved in connecting and registering with the server.
// After passing all connection/registration tests, add them to the client list
// and create their send/recieve channels.
func newClient(conn net.Conn, ch chan string, clients []Client) {
	newClient := &Client{
		conn:  conn,
		in:    make(chan string),
		out:   ch,
		suite: tests.NewTestSuite(),
		quit:  make(chan bool),
		list:  &clients,
	}

	// All Join tests must be passed before the client can move forward.
	go func() {
		newClient.runJoinTests()

		go clientsender(newClient)
		go clientreceiver(newClient)

		clients = append(clients, *newClient)
		fmt.Println("added to client list")
	}()
}

// Listen to a clients "in channel" and write anything sent over it to their
// connection.  If the quit bool is true, then close the connection and end
// the function.
func clientsender(client *Client) {
	for {
		select {
		case buf := <-client.in:
			client.conn.Write([]byte(buf))
			fmt.Println(buf)
		case <-client.quit:
			client.conn.Close()
			break
		}
	}
}

// Listen to the client's connection and if they send anything, push it to the client.out.
// If the client is quitting, close the connection and end the function.
func clientreceiver(client *Client) {
	buf := make([]byte, 2048)

	for client.Read(buf) {

		fmt.Println(string(buf))

		if bytes.EqualFold(buf, []byte("/quit")) {
			client.Close()
			break
		}

		send := client.nick + "> " + string(buf)
		client.out <- send
		for i := 0; i < 2048; i++ {
			buf[i] = 0x00
		}
	}

	client.out <- client.nick + " has left chat"
}
