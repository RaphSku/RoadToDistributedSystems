.PHONY: default
default: help

.PHONY: install_dependencies
## Install Swagger
install_dependencies:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: generate_private_public_key
## Generate a private and public key
generate_private_public_key:
	@go run ./utilities/pk.go

.PHONY: swagger
## Generate a new swagger spec
swagger:
	@swagger generate spec -o ./swagger.yml --scan-models

.PHONY: start_server
## Start server
start_server:
	@go run main.go

.PHONY: help
## Print this help
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
