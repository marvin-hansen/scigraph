# Make will use bash instead of sh
SHELL := /usr/bin/env bash
CC=clang # required by bazel
#ENV=LOCAL # Required by auto-config.

# GNU make man page
# http://www.gnu.org/software/make/manual/make.html

# For some strange reasons, intends & blanks shift in bash when calling 'make' so the formatting below should align intend at least on Bash on OSX.
.PHONY: help
help:
	@echo ' '
	@echo 'Setup: '
	@echo '    make check        	    	Checks whether all project requirements are present.'
	@echo '    make db-deploy        	Deploys & starts the DB container. Run just once to create. Then use docker container start/stop timescaledb.'
	@echo ' '
	@echo 'DB: '
	@echo '    make db-configure   	Configures the initial DB.Run again to reset & delete all data!'
	@echo '    make db-start   		starts the local timescale database (tsdb).'
	@echo '    make db-stop   		stops  the local timescale database (tsdb).'
	@echo ' '
	@echo 'Dev: '
	@echo '    make build   		Builds the code base incrementally (fast). Use for coding.'
	@echo '    make rebuild   		Rebuilds all dependencies & the code base (slow). Use after go mod changes. '
	@echo '    make stats   		Crunches & shows the latest project stats. '

#	@echo ' '
#	@echo 'Docker: '
#	@echo '    make build-smdb   		builds the docker image for service management & database service (SMDB).'


# "---------------------------------------------------------"
# Setup
# "---------------------------------------------------------"
.PHONY: check
check:
	@source scripts/setup/check_requirements.sh

.PHONY: db-deploy
db-deploy:
	@source  scripts/db/local/db_setup.sh


.PHONY: gen-keys
gen-keys:
	@source scripts/run/run_gen_keys.sh

# "---------------------------------------------------------"
# DB
# "---------------------------------------------------------"
.PHONY: db-configure
db-configure:
	@source  scripts/db/local/configure-local-db.sh

.PHONY: db-start
db-start:
	@source scripts/db/local/start_db.sh

.PHONY: db-stop
db-stop:
	@source scripts/db/local/stop_db.sh


# "---------------------------------------------------------"
# Development
# "---------------------------------------------------------"
.PHONY: build
build:
	@source scripts/dev/build_all.sh

.PHONY: rebuild
rebuild:
	@source scripts/dev/rebuild_all.sh

.PHONY: stats
stats:
	@source scripts/dev/project-stats.sh


# "---------------------------------------------------------"
# Services
# "---------------------------------------------------------"
