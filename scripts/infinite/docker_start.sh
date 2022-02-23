#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

command docker container start infinity

command docker exec infinity /app/restart.sh



