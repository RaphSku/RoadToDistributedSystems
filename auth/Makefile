# --- SHELL
SHELL := /usr/bin/env bash

# --- HELPERS
define check_tool_version
@if ! command -v $(1) >/dev/null 2>&1; then \
	echo "$(1) is not installed."; \
	echo "Minimum version required: $(3)"; \
	exit 1; \
fi; \
tool_version=$$($(2) | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+' | head -n1); \
if [ -z $$tool_version ]; then \
	echo "Could not determine version of $(1). Output of '$(2)' did not contain a version."; \
	exit 1; \
fi; \
min_version=$(3); \
max_version=$(4); \
if [ "$$(printf '%s\n' $$min_version $$tool_version | sort -V | head -n1)" != "$$min_version" ]; then \
	echo "$(1) version $$tool_version is too old. Please update to at least $$min_version."; \
	exit 1; \
elif [ -n "$$max_version" ] && \
     { [ "$$tool_version" = "$$max_version" ] || \
       [ "$$(printf '%s\n' $$tool_version $$max_version | sort -V | tail -n1)" != "$$max_version" ]; }; then \
	echo "$(1) version $$tool_version is too new. Please use a version lower than $$max_version."; \
	exit 1; \
else \
	if [ -n "$$max_version" ]; then \
		echo "$(1) version $$tool_version is OK (between $$min_version and < $$max_version)."; \
	else \
		echo "$(1) version $$tool_version is OK (>= $$min_version, < $$max_version)."; \
	fi; \
fi
endef

# --- VARIABLES
POSTGRES_USER := $(shell $(SHELL) -c "grep ^POSTGRES_USER= .env | cut -d '=' -f2-")

# --- TARGETS
.PHONY: default
default: help

.PHONY: preflight
preflight:
	$(call check_tool_version,docker,docker --version,28.0.4)
	$(call check_tool_version,docker compose,docker compose version,2.34.0)

.PHONY: start_psql_db
## Start Postgres container
start_psql_db: preflight
	docker compose -f postgres-compose.yaml up -d

.PHONY: take_down_psql_db
## Removes Postgres container
take_down_psql_db: preflight
	docker compose -f postgres-compose.yaml down --volumes

.PHONY: access_database
## Access Postgres database with psql
access_database: preflight
	echo $(POSTGRES_USER)
	docker compose -f postgres-compose.yaml exec -it postgres-auth psql -U $(POSTGRES_USER)

.PHONY: help
## Print this help screen
help:
	@echo "----------------------------------"
	@echo "Welcome to make! Enjoy the flight."
	@echo "Makefile - make [\033[38;5;154mtarget\033[0m]"
	@echo "----------------------------------"
	@echo
	@echo "Targets:"
	@awk '/^[a-zA-z\-_0-9%:\\]+/ { \
		description = match(descriptionLine, /^## (.*)/); \
		if (description) { \
			target = $$1; \
			description = substr(descriptionLine, RSTART + 3, RLENGTH); \
			gsub("\\\\", "", target); \
			gsub(":+$$", "", target); \
			printf "    \033[38;5;154m%-35s\033[0m %s\n", target, description; \
		} \
	} \
	{ descriptionLine = $$0 }' $(MAKEFILE_LIST)
	@printf "\n"
