#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

docker run -d --platform linux/x86_64 --name infinity -p 8190:8190 -p 8185:8185 infinity:latest
