## TCP Madness

This is a simple TCP server that handles message from a TCP client of a given format.

## Running the Go Server
1. `cd go`
2. `go build`
3. `./go`

The server is now up and running on `localhost:8080`

## Running the Client
1. `cd integration`
2. `./mac-distro`, or whatever distro corresponds with your OS

## Unanswered Questions
1. The integration tests end MUCH faster (10x faster at least) with some bogus logging. If I remove the extra logging, the tests are much slower. I'm not sure why that would be.
https://github.com/danbucholtz/tcp-madness/blob/integration-tests/go/tcp-server.go#L49

2. Am I using the mutex correctly? I lock and unlock around all read and writes of the data store
https://github.com/danbucholtz/tcp-madness/blob/integration-tests/go/data-access.go#L7

