.PHONY: default
default: help

.PHONY: preflight
preflight:
	@which go | test -n
	@go version

.PHONY: install_dependencies
## Install Go dependencies
install_dependencies:
	@go mod download

.PHONY: start
## Run code
start: preflight
	@go run main.go

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
