## TCP Madness

This project implements two identical-ish TCP servers. One uses the Golang stdlib. The other uses TypeScript and the Node.js stdlib.

## What Does This Do?

Each server launches and starts listening on port `8080` on `0.0.0.0`. The server handles requests that send messages of a given format, and return a given format. Ultimately, it implements an indexing service, where packages can be indexed, queried, and removed.

## Why Two Services?

Originally I implemented the server in Node.js and TypeScript, the tool I am most comfortable with. Once I had the server implemented and working correctly, I decided that I wanted to learn something new and dive into Golang. So I implemented the exact same service in Golang, as well as the same unit tests. Ultimately, it was a really good learning exercise and I'm very glad I did it.

## Building and Running the Go Server

Execute the following commands to run the Go server. It's assumed `golang` is installed on the system.

```
cd go // enter the go directory
go build // compile the go binary
./go // execute the go command
```

The Go server is now up and running on `localhost:8080` and ready for messages.

## Running Go Unit Tests

From the `go` directory, run `go test`.

## Building and Running the Node Server

The Node server is implemented in TypeScript, an excellent type system for JavaScript built by Microsoft. To run the node server, execute the following steps. It's assumed `node 8.x` and `npm 5.x` are installed. It's best to use the latest LTS release.

```
cd node // enter the node directory
npm install // install TypeScript and other build time dependencies
npm run start // run the start script from the npm scripts (package.json)
```
The `npm run start` command will compile the TypeScript, and start the server on `localhost:8080`

## Running Node Unit Tests

From the `node` directory, run `npm run test`. Since Node does not have a built-in testing utility, this project uses Jest, the most popular JavaScript test runner.

## Integration Tests

Since this is not a production app and time is a factor, this project only implements integration tests for the `mac os` platform. If you are using a different OS/platform than `mac os`, there are distros on the integration test harness, but you will have to manually start the server and evaluate the test harness results yourself.

To launch the fully automated integration tests for Mac, run the following commands from the project root.

Note: The `go` and `node` servers must be compiled/built before running the integration test script.

```
./integration/integration-test-go-mac.sh // for the Go server
./integration/integration-test-node-mac.sh // for the Node server
```

The integration suite launches a server on `localhost:8080`, and then executes the provided test harness for the `mac-os` platform against it. It then parses the test harness output for `All tests pass` and returns a `0` or `1` status code depending on if the tests pass or fail.

Since the script correctly fails with a non-zero status code, it would be appropriate to use in continuous integration environments.

## Docker Rational

It's been a few years since I've used Docker, so in the spirit of learning new things, I implemented Dockerfiles for both the `node` and `go` servers.

## Launching Server with Docker

In the `./scripts` directory, there are two bash scripts for building Docker images, and launching Docker containers for the images. Each container spins up a server listening on port `8080`, and then bash script maps `localhost:8080` to the Docker container's port `8080`.

To launch the `go` server as a Docker container, run:

```
./scripts/build-and-run-docker-go.sh
```

To launch the `node` server as a Docker container, run:

```
./scripts/build-and-run-docker-node.sh
```

At this point, the server is available on localhost, so any of the test harnesses in `./integration` can be run successfully against the server.

## Logging Details

Logging writes to `stdout` for now, with the intention that in "production" one would pipe the output to a file. In the case this application was ever deployed at robust scale, a more robust logging solution would be required. For all intents and purposes, piping to a file will be fine for this case.

There are two log levels. `warn` and `debug`. By default, the servers run at the `warn` setting. The log level can be changed by setting `LOGGING_LEVEL` environment variable to `debug`.

## Design Rational of the Go Server

The `go` server kicks off from `main.go`, where the logger is initialized and the TCP server kicks off. The logger reads for an environment variable to determine the setting, and then initializes the data it needs to efficiently log in a `log4j` fashion.

In `tcp-server.go`, the server is configured and started, and then when a new connection is accepted, a go routine executes for the request. In the spirit of separation of concerns, `tcp-server.go` knows only of the connection, and an unexpected error response. It delegates the details of processing a request to other parts of the solution. There is a `RequestStringtoResponseString` function in `tcp-server.go` that takes a vanilla string, and ultimately executes the validation and processing of the request string. It converts that data into the correct format for the response.

In `validate-command.go`, `ValidateCommand` is a function that takes a vanilla string, and performs a series of validation upon it to ensure that it is a valid request and meets a series of criteria. It took a little bit of time to figure out the validation and all of the rules, so there are extensive unit tests for this to ensure correctness. Ultimately this method returns a `Command *` or an `Error`, making it very easy for the consumer to work with.

In `process.command.go`, The `ProcessCommand` function takes the Command, and executes the necessary business logic against the `data-access.go` module. The business logic is pretty simple and easy to understand from reading the code, but there are unit tests as well to validate it.

In `data-access.go`, the application was kept simple and used an in-memory map data store. Since there are potentially `n` threads mutating/reading from the map at any given time, we use a read/write lock to ensure correctness and efficiency in accessing the data store.

In the `helpers.go` module, there are some basic utility functions and corresponding unit tests.

Overall, it's a pretty simple and easy-to-follow flow that focuses on separation of concerns and testability.

## Design Rational of the Node Server

The `node` server is virtually identical to the `go` server in rational and design with the exception of a few small details.

Since `node` does not allow for blocking, any request to the `data-store` is made asynchronous for the sake of A) being realistic and B) making it easier to move from the in-memory data store to a service or database or something.

Since `node` doesn't block and is "single threaded", there is no concept of a lock and there is no need or way to synchronize data access.

