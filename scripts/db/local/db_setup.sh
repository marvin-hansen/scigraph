# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
# set -o pipefail

command docker network create my-net

command docker run -d --name timescaledb --network my-net -p 5432:5432 -e POSTGRES_PASSWORD=2ee2e41ec0a7e6441e0038 timescale/timescaledb:latest-pg14
