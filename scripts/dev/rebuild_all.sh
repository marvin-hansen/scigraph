#
# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com
#

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

bazel build

# Convert mod dependencies into bazel dependencies
bazel run //:gazelle -- update-repos -from_file=go.mod

# Update all build files & dependencies
bazel run //:gazelle

# Regenerate all *.proto.pb file
# bazel run //proto:link

# Build all sources
bazel build //:build
