# --- HELPERS
define check_tool_version
@if ! command -v $(1) >/dev/null 2>&1; then \
	echo "$(1) is not installed."; \
	echo "Minimum version required: $(3)"; \
	echo "$(5)"; \
	echo "But the version should be lower than $(4)"; \
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
MIN_BUF_VERSION := 1.54.0
MAX_BUF_VERSION := 2.0.0
MIN_PROTOC_GEN_GO_VERSION := 1.36.6
MAX_PROTOC_GEN_GO_VERSION := 2.0.0
MIN_PROTOC_GEN_GO_GRPC_VERSION := 1.5.1
MAX_PROTOC_GEN_GO_GRPC_VERSION := 2.0.0

# --- TARGETS
.PHONY: default
default: help

.PHONY: preflight
preflight:
	@echo "Starting Preflight checks..."
	$(call check_tool_version, \
		go, \
		go version,$(MIN_GO_VERSION),$(MAX_GO_VERSION), \
		"You seem to be missing Go\, you can download Go here: https://go.dev/dl/")
	$(call check_tool_version, \
		buf, \
		buf --version,$(MIN_BUF_VERSION),$(MAX_BUF_VERSION), \
		"You seem to be missing buf\, you can get it here: https://buf.build/docs/cli/installation/")
	$(call check_tool_verion, \
		protoc-gen-go, \
		protoc-gen-go --version,$(MIN_PROTOC_GEN_GO_VERSION),$(MAX_PROTOC_GEN_GO_VERSION), \
		"You seem to be missing protoc-gen-go\, install with the following cmd:" \
		"go install google.golang.org/protobuf/cmd/protoc-gen-go@latest")
	$(call check_tool_version, \
		protoc-gen-go-grpc, \
		protoc-gen-go-grpc --version,$(MIN_PROTOC_GEN_GO_GRPC_VERSION),$(MAX_PROTOC_GEN_GO_GRPC_VERSION), \
		"You seem to be missing protoc-gen-go-grpc\, install with the following cmd:" \
		"go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest")
	@echo "Preflight checks succeeded..."

.PHONY: proto
proto: preflight
	buf generate
	cp -r ./gen/itemstore/proto/* ./client/itemstore/
	cp -r ./gen/itemstore/proto/* ./server/itemstore/

.PHONY: start_client
## Start gRPC client application
start_client: proto
	cd client && go run main.go

.PHONY: start_server
## Start gRPC server application
start_server: proto
	cd server && go run main.go

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
