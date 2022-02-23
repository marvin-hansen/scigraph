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
	@echo '    make check        	    	  Checks whether all project requirements are present.'
	@echo '    make db-deploy        	  Deploys & starts the DB container. Run just once to create'
	@echo '    make graph-build-image    	  Builds a docker image containing infinite graph.'
	@echo '    make graph-create-container   Creates an infinite graph cocker container that can be started & stopped.'
	@echo '    make graph-delete-container   Deletes the infinite graph cocker container that can be started & stopped.'
	@echo ' '
	@echo 'GRAPH: '
	@echo '    make graph-start   		Starts the local infinite graph.'
	@echo '    make graph-stop   		Stops  the local infinite graph.'

	@echo 'DB: '
	@echo '    make db-configure   	Configures the initial DB.Run again to reset & delete all data!'
	@echo '    make db-start   		Starts the local timescale database (tsdb).'
	@echo '    make db-stop   		Stops  the local timescale database (tsdb).'
	@echo ' '
	@echo 'Dev: '
	@echo '    make build   		Builds the code base incrementally (fast). Use for coding.'
	@echo '    make rebuild   		Rebuilds all dependencies & the code base (slow). Use after go mod changes. '
	@echo '    make stats   		Crunches & shows the latest project stats. '

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
# INFINITE GRAPH
# "---------------------------------------------------------"
.PHONY: graph-build-image
graph-build-image:
	@source scripts/infinite/docker_build.sh

.PHONY: graph-create-container
graph-create-container:
	@source scripts/infinite/docker_container_run.sh

.PHONY: graph-delete-container
graph-delete-container:
	@source scripts/infinite/docker_container_remove.sh

.PHONY: graph-start
graph-start:
	@source scripts/infinite/docker_start.sh

.PHONY: graph-stop
graph-stop:
	@source scripts/infinite/docker_stop.sh

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
