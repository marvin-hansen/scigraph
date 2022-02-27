#!/bin/bash
#!/usr/bin/env bash

echo "Run script:"

echo " * Start Lock Server"
command objy startlockserver  # >/dev/null

command sleep 1  # Waits 1 second so that lockserver is online.

echo " * Start Studio Server"
command objy startstudioserver

command sleep 1  # Wait a bit to let studio server start

# Need to configure DB before starting API server to serve DB over API
# You can start the Objectivity REST server by providing a single argument, which is the path to the bootfile of a federated database.

# echo " * Start API Server"
# command objy StartRESTServer -fdAlias kb -configFile restApiConfig.xml

# Check if API service is running.
# objy CheckRESTServer

# This kees the container running after the script completed,
# but will exit immediately on Ctr-C or a docker stop signal
exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"
