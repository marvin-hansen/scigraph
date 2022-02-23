#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

command docker build --platform linux/x86_64 -t infinity:latest .