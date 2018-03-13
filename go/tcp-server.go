package main

import (
	"bufio"
	"net"
	"os"
	"time"
)

// HOST IP Address or name
const HOST = "0.0.0.0"

// PORT to listen on
const PORT = "8080"

// TYPE of network packets (TCP)
const TYPE = "tcp"

func startServer() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		Warnf("Server failed to start listening: ", err.Error())
		os.Exit(1)
	}

	// sweet, we're active and listenin'
	defer listener.Close()
	Warn("Listening on " + HOST + ":" + PORT)

	for {
		connection, err := listener.Accept()
		if err != nil {
			Warnf("Error Accepting Connecting", err.Error())
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
			// if err.Error() != "EOF" {
			// fmt.Printf("__%s__\n", err.Error())
			// }
			Warnf("Unexpected error when reading from request: ", err)
			conn.Write([]byte("ERROR\n"))
		} else if len(message) > 0 {
			Debugf("Request received %s", message)
			response := RequestStringtoResponseString(message)
			end := makeTimestamp()
			millis := end - start
			Debugf("Request Completed, took %d millis", millis)
			conn.Write([]byte(response))
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// RequestStringtoResponseString takes a raw string command from a TCP request and processes it, created a string to return as a response
func RequestStringtoResponseString(requestCommand string) string {
	// validate the request
	command, validationError := ValidateCommand(requestCommand)
	if validationError != nil {
		Debugf("Invalid Request: ", validationError)
		return "ERROR\n"
	}
	result, err := processCommand(command)
	if err != nil {
		Debugf("Error occurred when processing request: ", err)
		return "ERROR\n"
	}
	if result {
		Debug("Request was processed and operation was successful: ")
		return "OK\n"
	}
	Debug("Request was processed and operation was not successful: ")
	return "FAIL\n"
}
