#!/bin/sh

# Define the federation name.
FDNAME=example2

# Define the federation data directory.
DATA_DIR=data

# Define the name and path for the federation's configuration file.
BOOT_FILE=${FDNAME}.boot
BOOT_PATH=${DATA_DIR}/${BOOT_FILE}

if ! objy checkls -notitle; then 
  echo Please start the lock server before running this script.
  echo To start the lockserver run the following command from a user account that has admin privileges:
  echo objy startlockserver
  exit 1
fi

# Create a Data directory
if [ ! -d ${DATA_DIR} ]; then
  mkdir -p ${DATA_DIR}
fi

# Cleanly remove a previous run of this script.
if [ -e ${BOOT_PATH}  ]; then
  objy installfd -notitle -bootfile ${BOOT_PATH}
  objy cleanupfd -notitle -bootfile ${BOOT_PATH} -local
  objy deletefd  -notitle -bootfile ${BOOT_PATH}
fi

# Create the federation.
# A federation is composed of a configuration (*.boot) and data (*.fdb) file.
objy createFD -fdName ${FDNAME} -notitle -fddirp ./${DATA_DIR} || exit 1

# Execute a Do script to create the federation's schema.
objy do -infile createSchema.do -notitle -bootfile ${BOOT_PATH} || exit 1

# Create an index named Person_id for the person id attribute.
objy do -infile createData.do -notitle -bootfile ${BOOT_PATH} || exit 1

# Create an index named Phone_number for the phone number attribute.
objy addindex -name Person_id -class Person -attribute id -notitle -boot ${BOOT_PATH} || exit 1

# Populate the database with example data via a Do script.
# The script reads in data from CSV files.
objy addIndex -name Phone_number -class Phone -attribute phoneNumber -notitle -boot ${BOOT_PATH} || exit 1

# Query the database via a Do script and store the results in a file.
objy do -infile runQuery.do -outfile queryResults.txt -notitle -bootfile ${BOOT_PATH} || exit 1

cat queryResults.txt
