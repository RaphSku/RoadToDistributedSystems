start_server:
	@go run main.go

install_deps:
	go install github.com/cloudflare/cfssl/cmd/...@latest

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
