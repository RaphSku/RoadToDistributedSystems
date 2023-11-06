.PHONY: install_dependencies generate_private_key init_swag start_server generate_swag

install_dependencies:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate_private_key:
	@go run ./utilities/pk.go

swagger:
	@swagger generate spec -o ./swagger.yml --scan-models

start_server:
	@go run main.go