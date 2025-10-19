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
MIN_GO_VERSION=1.24.4
MAX_GO_VERSION=1.25.0

# --- TARGETS
default: help

.PHONY: preflight
preflight:
	$(call check_tool_version,go,go version,$(MIN_GO_VERSION),$(MAX_GO_VERSION))
	which go
	which cfssl || go install github.com/cloudflare/cfssl/cmd/...@latest

.PHONY: gencert
## Generate certificates
gencert: preflight
	cfssl gencert \
		-initca certs-config/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=certs-config/ca-config.json \
		-profile=server \
		certs-config/server-csr.json | cfssljson -bare server

	mv *.pem *.csr certs/

.PHONY: start_server
## Start server
start_server: gencert
	@echo "Starting Server..."
	@go run main.go
	@echo "Server finished..."

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
