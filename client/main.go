package main

import (
	"bufio"
	"fmt"
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
	welcome_msg := fmt.Sprintf("Welcome %s, to the chat and say hi..", username)
	fmt.Println(welcome_msg)

}
