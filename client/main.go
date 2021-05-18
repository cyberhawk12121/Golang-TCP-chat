package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", ":8080")
	logFatal(err)

	defer connection.Close()
	// we need to listen to what the user enters at the terminal
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	logFatal(err)
	username = strings.Trim(username, "\r\n")

	//welcome message
	welcome_msg := fmt.Sprintf("Welcome %s, to the chat and say hi..\n", username)
	fmt.Println(welcome_msg)

	// Read other's messages
	go read(connection)
	// Write message from the terminal
	write(connection, username)
}

func read(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			conn.Close()
			fmt.Println("Connection closed")
			os.Exit(0)
		}
		fmt.Println("--> ", message)
	}
}

func write(conn net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// string formating
		message = fmt.Sprintf("%s: -%s\n", username, strings.Trim(message, "\r\n"))
		conn.Write([]byte(message))
	}
}
