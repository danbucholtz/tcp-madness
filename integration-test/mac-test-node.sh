#!/bin/bash
node node/dist/main.js &
serverPid=$!
echo "Starting up the Node implementation of the TCP server with PID $serverPid"
sleep 5
echo "Starting Mac (darwin) test harness"
./integration-test/mac-distro -concurrency 100 &> ./integration-test/node-server-integration-results.txt
blah=$(cat ./integration-test/node-server-integration-results.txt | grep "All tests passed")
retVal=$?
if [ ! $? -eq 0 ]; then
    echo "The Integration Tests Failed"
fi
echo "Shutting down server"
kill $serverPid
echo "Deleting tmp files"
rm -rf ./integration-test/node-server-integration-results.txt
echo "Exiting with return code $retVal"
exit $retVal

