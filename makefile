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
	@echo '    make graph-build-image    	  Builds a docker image containing infinite graph.'
	@echo '    make graph-create-container   Creates an infinite graph cocker container that can be started & stopped.'
	@echo '    make graph-delete-container   Deletes the infinite graph cocker container that can be started & stopped.'
	@echo '    make graph-connect-container   Deletes the infinite graph cocker container that can be started & stopped.'
	@echo ' '
	@echo 'GRAPH: '
	@echo '    make graph-start   		Starts the local infinite graph.'
	@echo '    make graph-stop   		Stops  the local infinite graph.'
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

# "---------------------------------------------------------"
# INFINITE GRAPH
# "---------------------------------------------------------"
.PHONY: graph-build-image
graph-build-image:
	@source scripts/infinite/docker_build.sh

.PHONY: graph-create-container
graph-create-container:
	@source scripts/infinite/docker_run.sh

.PHONY: graph-delete-container
graph-delete-container:
	@source scripts/infinite/docker_remove.sh

.PHONY: graph-connect-container
graph-connect-container:
	@source scripts/infinite/docker_exec.sh

.PHONY: graph-start
graph-start:
	@source scripts/infinite/docker_start.sh

.PHONY: graph-stop
graph-stop:
	@source scripts/infinite/docker_stop.sh

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
