package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"

	hclog "github.com/hashicorp/go-hclog"
)

type ConfigManager struct {
	logger hclog.Logger
}

func (cm *ConfigManager) init(logger hclog.Logger) {
	cm.logger = logger
}

func (cm *ConfigManager) readCAFile(path string) []byte {
	caFile, err := os.ReadFile("certs/ca.pem")
	if err != nil {
		cm.logger.Error(fmt.Sprintf("CA Certificate could not be read due to:%s", err))
		os.Exit(1)
	}

	return caFile
}

func main() {
	logger := hclog.Default()

	cm := ConfigManager{}
	cm.init(logger)

	caFile := cm.readCAFile("ca.pem")

	ca := x509.NewCertPool()
	ok := ca.AppendCertsFromPEM([]byte(caFile))
	if !ok {
		cm.logger.Error("CA Certificate could not be appended to the Cert Pool")
		os.Exit(1)
	}

	config := &tls.Config{
		RootCAs: ca,
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:9090/")
	if err != nil {
		cm.logger.Error(fmt.Sprintf("GET request failed due to:%s", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			cm.logger.Error(fmt.Sprintf("Body could not be read due to:%s", err))
			os.Exit(1)
		}
		body := string(bodyBytes)
		println(body)
	}
}
