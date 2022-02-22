# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

docker run --platform linux/x86_64 --name infity --net my-net -p 8190:8190 infity:latest
