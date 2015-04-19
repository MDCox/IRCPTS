package main

import (
	"./tests"
	"bytes"
	"net"
	"strings"
)

type Client struct {
	nick  string      // Client name
	conn  net.Conn    // Client connection
	in    chan string // Channel that sends messages to client
	out   chan string // Channel that sends messages to others
	suite tests.Suite // Client's test suite

	quit chan bool // quit channel for all goroutines
	list *[]Client // reference to list
}

func (c *Client) runTest(t *tests.Test, msg string) bool {
	if t.Eval(msg) {
		c.suite.Passed = append(c.suite.Passed, *t)
	} else {
		c.suite.Failed = append(c.suite.Failed, *t)
	}

	return t.Result
}

// Run tests involved in client join
func (c *Client) runJoinTests() {
	buf := make([]byte, 1024)

	joinTests := c.suite.Tests[0].Tests

	for {
		if joinTests[1].Result && joinTests[2].Result {
			return
		}

		c.Read(buf)

		input := strings.Split(string(buf), "\r\n")
		for _, line := range input {
			switch line[0:4] {
			case "PASS":
				c.runTest(&joinTests[0], line)
			case "NICK":
				c.nick = line[4:]
				c.runTest(&joinTests[1], line)
			case "USER":
				c.runTest(&joinTests[2], line)
			}
		}
	}
}

// Read from connection and return true if ok
func (c Client) Read(buf []byte) bool {
	_, err := c.conn.Read(buf)
	if err != nil {
		c.Close()
		return false
	}

	return true
}

// Close connection and remove from client list
func (c *Client) Close() {
	c.quit <- true
	c.conn.Close()
	//c.deleteFromList()
}

// Compare two clients
func (c *Client) Equal(cl *Client) bool {
	if bytes.EqualFold([]byte(c.nick), []byte(cl.nick)) {
		if c.conn == cl.conn {
			return true
		}
	}
	return false
}
