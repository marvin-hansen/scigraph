# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

command docker container start infinity

command docker exec infinity /app/restart.sh



