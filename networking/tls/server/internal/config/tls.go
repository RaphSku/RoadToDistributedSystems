package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

type TLSConfig struct {
	CAFile        string
	ServerAddress string
	Server        bool
	CertFile      string
	KeyFile       string
}

func SetupTLSConfig(config TLSConfig) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	if config.CertFile != "" && config.KeyFile != "" {
		tlsConfig.Certificates = make([]tls.Certificate, 1)
		keypair, err := tls.LoadX509KeyPair(
			config.CertFile,
			config.KeyFile,
		)
		tlsConfig.Certificates[0] = keypair
		if err != nil {
			return nil, err
		}
	}

	if config.CAFile != "" {
		caFile, err := os.ReadFile(config.CAFile)
		if err != nil {
			return nil, err
		}

		ca := x509.NewCertPool()
		ok := ca.AppendCertsFromPEM([]byte(caFile))
		if !ok {
			return nil, fmt.Errorf("Failed to parse root certificate: %q", config.CAFile)
		}

		if config.Server {
			tlsConfig.ClientCAs = ca
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		} else {
			tlsConfig.RootCAs = ca
		}
		tlsConfig.ServerName = config.ServerAddress
	}

	return tlsConfig, nil
}
