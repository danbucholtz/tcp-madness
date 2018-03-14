#!/bin/bash
cd go
echo "Building the Docker image"
docker build -t go-tcp-madness .
echo "Starting up the container"
docker run -p 8080:8080 go-tcp-madness