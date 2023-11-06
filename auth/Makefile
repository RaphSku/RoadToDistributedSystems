.PHONY: default preflight start_psql_db help

default: help

preflight:
	@which docker | test -n
	@docker --version
	@docker compose version

## Start docker container with Postgres
start_psql_db: preflight
	@docker compose -f postgres-compose.yaml up

## Print this help screen
help:
	@printf -- "-------------------------------\n"
	@printf "Makefile - make [target]\n"
	@printf -- "-------------------------------\n\n"
	@printf "Targets:\n\n"
	@awk '/^[a-zA-Z\-_0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
		helpCommand = $$1; \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
		printf "  \x1b[91;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"