package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		fmt.Println("New Request received!")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("ERROR\n"))
			// conn.Close()
		} else {
			response := RequestStringtoResponseString(message)
			conn.Write([]byte(response))
			// conn.Close()
		}
	}
}

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

/*func processRequest(conn net.Conn) {
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading to buffer: ", err.Error())
	}

	untrimmed := string(buf)
	rawCommandString := strings.Trim(untrimmed, " ")

	fmt.Printf("About to process %s \n", rawCommandString)
	command, validationError := ValidateCommand(rawCommandString)
	if validationError != nil {
		fmt.Println(validationError)
		fmt.Printf("Validation error\n")
		conn.Write([]byte("ERROR\n"))
		conn.Close()
		// return
	}

	fmt.Println(command)
	result, err := processCommand(command)
	if err != nil {
		fmt.Printf("Error processing the command\n")
		conn.Write([]byte("ERROR\n"))
		conn.Close()
		// return
	} else if result {
		fmt.Println("Sending back OK")
		conn.Write([]byte("OK\n"))
		conn.Close()
		// return
	} else {
		fmt.Println("Sending back FAIL")
		conn.Write([]byte("FAIL\n"))
		conn.Close()
	}
}
*/
