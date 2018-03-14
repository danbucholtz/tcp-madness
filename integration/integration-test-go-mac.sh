#!/bin/bash
./go/go &
serverPid=$!
echo "Starting up the Go implementation of the TCP server with PID $serverPid"
sleep 5
echo "Starting Mac (darwin) test harness"
./integration/mac-distro -concurrency 100 &> ./integration/test-results-go.txt
blah=$(cat ./integration/test-results-go.txt | grep "All tests passed")
retVal=$?
if [ ! $? -eq 0 ]; then
    echo "The Integration Tests Failed"
fi
echo "Shutting down server"
kill $serverPid
echo "Deleting tmp files"
rm -rf ./integration/test-results-go.txt
echo "Exiting with return code $retVal"
exit $retVal

