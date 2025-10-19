# --- HELPERS
define check_tool_version
@if ! command -v $(1) >/dev/null 2>&1; then \
	echo "$(1) is not installed.\nMinimum version required: $(2)"; \
	exit 1; \
fi; \
tool_version=$$($(2) | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+' | head -n1); \
if [[ -z $$tool_version ]]; then \
	echo "Could not determine version of $(1). Output of '$(2)' did not contain a version."; \
	exit 1; \
fi; \
required_version=$(3); \
if [[ "$$(printf '%s\n' $$required_version $$tool_version | sort -V | head -n1)" != "$$required_version" ]]; then \
	echo "$(1) version $$tool_version is too old. Update please to at least $$required_version."; \
	exit 1; \
else \
	echo "$(1) version $$tool_version is OK."; \
fi
endef

# --- VARIABLES
SWAGGER_VERSION := 0.32.3

# --- TARGETS
.PHONY: default
default: help

.PHONY: preflight
preflight:
	$(call check_tool_version,swagger,swagger version,$(SWAGGER_VERSION))

.PHONY: install_dependencies
## Install swagger
install_dependencies:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@v$(SWAGGER_VERSION)

.PHONY: swagger
## Generate a new swagger spec
swagger: preflight
	swagger generate spec -o ./swagger.yml --scan-models

.PHONY: start_server
## Start server
start_server:
	go run main.go

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
