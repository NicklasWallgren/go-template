#####################################################
# Make variables			 						#
#####################################################

# The default environment file
ENVIRONMENT_FILE=$(shell pwd)/.env

# Available docker containers
CONTAINERS=app mariadb rabbitmq

# Default environment variables
MYSQL_ROOT_PASSWORD ?= secret
MYSQL_DATABASE ?= go_template

#####################################################
# RUNTIME TARGETS			 						#
#####################################################

default: run

# Start the containers
run: prerequisite build

# Start individual container
start: prerequisite valid-container
	- docker-compose -f docker-compose.yml up --build $(filter-out $@,$(MAKECMDGOALS))

# Start individual container in the background
silent-start: prerequisite valid-container
	- docker-compose -f docker-compose.yml up -d --build $(filter-out $@,$(MAKECMDGOALS))

# Quick start the containers, no preparation will be done
quick-start: prerequisite build-containers

# Quick start of dependencies
start-dependencies:
	- docker-compose -f docker-compose.yml up -d --build mariadb rabbitmq

# Apply prerequisites such as database schema
rollup: database-truncate-and-apply
	- @echo "TODO";

# Stop individual container
stop: prerequisite valid-container
	- docker-compose -f docker-compose.yml stop $(filter-out $@,$(MAKECMDGOALS))

# Halts the docker containers
halt: prerequisite
	- docker-compose -f docker-compose.yml kill

#####################################################
# SETUP AND BUILD TARGETS			 				#
#####################################################

# Build and prepare the docker containers and the project
build: prerequisite build-containers

# Build and launch the containers
build-containers:
	- docker-compose -f docker-compose.yml up --build

# Remove the docker containers
clean: prerequisite prompt-continue
	# Remove the docker containers, networks and volumes
	- docker-compose -f docker-compose.yml rm -svf
	- docker-compose -f docker-compose.yml down --rmi all -v --remove-orphans
	- rm -rf app/docker/data

# Echos the container status
status: prerequisite
	- docker-compose -f docker-compose.yml ps

# Drop and creates a new database
database-setup:
	- mysql -u root -p$(MYSQL_ROOT_PASSWORD) -h 127.0.0.1 -e "DROP DATABASE IF EXISTS ${MYSQL_DATABASE}; CREATE DATABASE ${MYSQL_DATABASE};"
	# TODO, apply migration and seed

# Drop and creates a new test database
database-test-setup:
	- mysql -u root -p$(MYSQL_ROOT_PASSWORD) -h 127.0.0.1 -e "DROP DATABASE IF EXISTS ${MYSQL_DATABASE}_test; CREATE DATABASE ${MYSQL_DATABASE}_test;"

#####################################################
# BASH CLI TARGETS			 						#
#####################################################

# Opens a bash prompt to the php cli container
bash-app: prerequisite
	- docker-compose -f docker-compose.yml exec --env COLUMNS=`tput cols` --env LINES=`tput lines` app bash

# Opens a bash prompt to the php fpm container
bash-mariadb: prerequisite
	- docker-compose -f docker-compose.yml exec --env COLUMNS=`tput cols` --env LINES=`tput lines` mariadb bash

# Opens the rabbitmq bash cli
bash-rabbitmq: prerequisite
	- docker-compose -f docker-compose.yml exec --env COLUMNS=`tput cols` --env LINES=`tput lines` rabbitmq bash

#####################################################
# APPLICATION CLI TARGETS			 				#
#####################################################

# Opens the mysql cli
cli-mariadb:
	- docker-compose -f docker-compose.yml exec mysql mysql -u root -p$(MYSQL_ROOT_PASSWORD)

#####################################################
# TEST TARGETS			 						    #
#####################################################





#####################################################
# INTERNAL TARGETS			 						#
#####################################################

# Validates the prerequisites such as environment variables
prerequisite: check-environment
	- @echo "pwd "$(shell pwd)
-include .env
export ENV_FILE = $(ENVIRONMENT_FILE)

# Validates the environment variables
check-environment:
	@echo "Validating the environment";

# Check whether the docker binary is available
ifeq (, $(shell which docker-compose))
	$(error "No docker-compose in $(PATH), consider installing docker")
endif

# Check whether the mysql-cli binary is available
ifeq (, $(shell which mysql))
	$(error "No mysql-cli in $(PATH), consider installing mysql-cli")
endif

# Validates the containers
valid-container:
ifeq ($(filter $(filter-out $@,$(MAKECMDGOALS)),$(CONTAINERS)),)
	$(error Invalid container provided "$(filter-out $@,$(MAKECMDGOALS))")
endif

# Prompt to continue
prompt-continue:
	@while [ -z "$$CONTINUE" ]; do \
		read -r -p "Would you like to continue? [y]" CONTINUE; \
	done ; \
	if [ ! $$CONTINUE == "y" ]; then \
        echo "Exiting." ; \
        exit 1 ; \
    fi

%:
	@: