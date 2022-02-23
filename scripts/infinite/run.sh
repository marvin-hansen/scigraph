#!/bin/bash
#!/usr/bin/env bash

echo "Run script:"

echo " * Start Lock Server"
command ./bin/objy startlockserver  # >/dev/null

command sleep 1  # Waits 1 second so that lockserver is online.

echo " * Start Studio Server"
command ./bin/objy startstudioserver  # >/dev/null

command sleep 1  # Wait a bit to let studio server start

# You can start the Objectivity REST server by providing a single argument, which is the path to the bootfile of a federated database.
#  Alternatively, you can supply a configuration file with aliases and paths for multiple federated databases.
# Using a configuration file lets you send REST commands to different federated databases through the use of a URL parameter;
# command ./bin/objy -configFile config.xml

# This kees the container running after the script completed,
# but will exit immediately on Ctr-C or a docker stop signal
exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"