package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// Host IP Address or name
const Host = "0.0.0.0"

// Port to listen on
const Port = "8080"

// Type of network packets (TCP)
const Type = "tcp"

func startServer() {
	listener, err := net.Listen(Type, Host+":"+Port)
	if err != nil {
		fmt.Println("Server failed to start listening: ", err.Error())
		os.Exit(1)
	}

	// sweet, we're active and chillin'
	defer listener.Close()
	fmt.Println("Listening on " + Host + ":" + Port)

	for {
		// listen for connections
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go processRequest(connection)
	}
}

func processRequest(conn net.Conn) {
	for {
		start := makeTimestamp()
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("ERROR\n"))
			// conn.Close()
		} else if len(message) > 0 {
			fmt.Printf("Request received %s\n", message)
			response := RequestStringtoResponseString(message)
			end := makeTimestamp()
			millis := end - start
			fmt.Printf("Request complete in %d millis\n", millis)
			conn.Write([]byte(response))
			// conn.Close()
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// RequestStringtoResponseString takes a raw string command from a TCP request and processes it, created a string to return as a response
func RequestStringtoResponseString(requestCommand string) string {
	command, validationError := ValidateCommand(requestCommand)
	if validationError != nil {
		return "ERROR\n"
	}
	result, err := processCommand(command)
	if err != nil {
		return "ERROR\n"
	}
	if result {
		return "OK\n"
	}
	return "FAIL\n"
}
