# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

# bin/bash

command ./bin/objy startlockserver >/dev/null

command sleep 1  # Waits 1 second so that lockserver is online.

command ./bin/objy startstudioserver

command sleep 1  # Wait a bit to let studio server start

# You can start the Objectivity REST server by providing a single argument, which is the path to the bootfile of a federated database.
#  Alternatively, you can supply a configuration file with aliases and paths for multiple federated databases.
# Using a configuration file lets you send REST commands to different federated databases through the use of a URL parameter;
# command ./bin/objy -configFile config.xml
