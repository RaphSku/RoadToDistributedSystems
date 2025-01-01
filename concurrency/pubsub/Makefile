.PHONY: default
default: help

.PHONY: start-producer
## Start the Producer
start-producer:
	cd producer && go run main.go

.PHONY: start-queue
## Start the Queue Server
start-queue:
	cd queue && go run main.go

.PHONY: start-consumer
## Start the Consumer
start-consumer:
	cd consumer && go run main.go

.PHONY: produce-messages
## Produce messages
produce-messages:
	cd bin && chmod +x produce-messages.sh
	cd bin && ./produce-messages.sh

.PHONY: get-messages
## Get messages
get-messages:
	curl -X GET http://localhost:10001/api/v1/consumer/messages

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
