#!/bin/bash
#!/usr/bin/env bash

# We have to restart the studio server after docker stop & start
echo " * Start Studio Server"
command ./bin/objy startstudioserver  # >/dev/null

command sleep 1  # Wait a bit to let studio server start
