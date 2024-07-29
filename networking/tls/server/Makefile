default: help

.PHONY: start_server
## Start server
start_server:
	@go run main.go

.PHONY: install_deps
## Install necessary dependencies
install_deps:
	go install github.com/cloudflare/cfssl/cmd/...@latest

.PHONY: gencert
## Generate certificates
gencert:
	cfssl gencert \
		-initca certs-config/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=certs-config/ca-config.json \
		-profile=server \
		certs-config/server-csr.json | cfssljson -bare server

	mv *.pem *.csr certs/

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
