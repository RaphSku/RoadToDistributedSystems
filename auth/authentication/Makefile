.PHONY: default install_dependencies generate_private_key init_swag start_server generate_swag

default: help

install_dependencies:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate_private_key:
	@go run ./utilities/pk.go

swagger:
	@swagger generate spec -o ./swagger.yml --scan-models

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