.PHONY: default install_dependencies generate_private_key init_swag start_server generate_swag

default: help

## Install dependencies like Swagger
install_dependencies:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

## Generate a private and public key
generate_private_public_key:
	@go run ./utilities/pk.go

## Generate a new swagger spec
swagger:
	@swagger generate spec -o ./swagger.yml --scan-models

## Start server
start_server:
	@go run main.go

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