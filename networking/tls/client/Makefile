default: help

.PHONY: test_curl
## Test TLS via curl
test_curl:
	curl --cacert ca.pem -X GET https://localhost:9090

.PHONY: fetch_ca_certificate
## Copy the necessary certificates over from the certs folder
fetch_ca_certificate:
	@cp ../server/certs/ca.pem ./certs/ca.pem

.PHONY: run_client
## Run client
run_client:
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
