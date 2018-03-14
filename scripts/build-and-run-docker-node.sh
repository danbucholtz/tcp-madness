#!/bin/bash
cd node
echo "Building the Docker image"
docker build -t node-tcp-madness .
echo "Starting up the container"
docker run -p 8080:8080 node-tcp-madness