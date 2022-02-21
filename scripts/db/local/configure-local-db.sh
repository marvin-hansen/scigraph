#
# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com
#

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Password is not used in this script as its exec's into the container and runs as authorized user on localhost inside the container.
# When the docker default password changes, this variable needs an update.
# Docker default password set in scripts/db/local/db_setup
# export PGPASSWORD=''
DB_HOST='timescaledb' # name of the docker container
DB_USER='postgres'

# Delete everything, if exists, to start anew from scratch.
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP DATABASE IF EXISTS configdb;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP DATABASE IF EXISTS tradedb;"

# Delete all users
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS configuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS tradeuser;"

command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS amxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS cmxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS smxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS imxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS tmxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS ftxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS vmxuser;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "DROP USER IF EXISTS vmxuser;"

# Create DB users
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER configuser with encrypted password 'configtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER tradeuser with encrypted password 'tradetest';"

# Create new user with password. Must match default local config in cmdb/config/smx_db.go
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER cmxuser with encrypted password 'cmxtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER smxuser with encrypted password 'smxtest';"

command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER amxuser with encrypted password 'amxtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER imxuser with encrypted password 'imxtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER tmxuser with encrypted password 'tmxtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER ftxuser with encrypted password 'ftxtest';"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE USER vmxuser with encrypted password 'vmxtest';"

# create new DB instances
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE DATABASE configdb WITH OWNER configuser TABLESPACE pg_default;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -c "CREATE DATABASE tradedb WITH OWNER tradeuser TABLESPACE pg_default;"

# Create DB schema for each instance & link it to DB
# configDB schemas
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "configdb" -c "CREATE SCHEMA IF NOT EXISTS amdb AUTHORIZATION amxuser;;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "configdb" -c "CREATE SCHEMA IF NOT EXISTS cmdb AUTHORIZATION cmxuser;;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "configdb" -c "CREATE SCHEMA IF NOT EXISTS smdb AUTHORIZATION smxuser;;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "configdb" -c "CREATE SCHEMA IF NOT EXISTS imdb AUTHORIZATION imxuser;;"

# tradeDB schemas
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "tradedb" -c "CREATE SCHEMA IF NOT EXISTS tmx  AUTHORIZATION tmxuser;;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "tradedb" -c "CREATE SCHEMA IF NOT EXISTS ftx  AUTHORIZATION ftxuser;;"
command docker exec -it "$DB_HOST" psql -U "$DB_USER" -h localhost -d "tradedb" -c "CREATE SCHEMA IF NOT EXISTS vmx  AUTHORIZATION vmxuser;;"
