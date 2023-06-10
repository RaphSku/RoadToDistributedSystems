
.PHONY: test_curl

test_curl:
	curl --cacert ca.pem -X GET https://localhost:9090

fetch_ca_certificate:
	@cp ../server/certs/ca.pem ./certs/ca.pem

run_client:
	@go run main.go