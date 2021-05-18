package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// variables
var (
	openConnections = make(map[net.Conn]bool) // This map is there to track all open connections. It's a key-value pair, key=connection, value=Boolean
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func broadcastMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// loop over all the connections
		// send message to those connections
		//except the connections that send the messages
		for item := range openConnections {
			if item != conn {
				item.Write([]byte(message))
			}
		}
	}
	deadConnection <- conn
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	logFatal(err)

	// At the end of the program close the connection
	defer ln.Close()

	go func() { // it's a goroutine
		// Whenever we get a request/connection from a new client it'll handle it
		for {
			conn, err := ln.Accept() //	Accept() returns a new connection
			logFatal(err)

			openConnections[conn] = true
			// To make it available outside the scope of this routine
			// we need to pass it to a channel
			newConnection <- conn
		}
	}()

	connection := <-newConnection
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	logFatal(err)
	fmt.Println(message)

	for {
		select {
		case conn := <-newConnection:
			//Invoke Broadcast()
			go broadcastMessage(conn)
		case conn := <-deadConnection:
			// remove the connection from the map
			for item := range openConnections {
				if item == conn {
					break
				}
			}
			delete(openConnections, conn)
		}
	}

}
