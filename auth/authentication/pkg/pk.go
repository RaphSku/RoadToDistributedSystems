package pkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"os"
	"path"
)

func GetPrivatePublicKey(pathToPEMs string) (string, string) {
	if _, err := os.Stat(path.Join(pathToPEMs, "pk.pem")); errors.Is(err, os.ErrNotExist) {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			println("Private and Public Key could not be generated due to:", err)
			os.Exit(1)
		}
		pemEncodedPrivateKey, pemEncodedPublicKey := Encode(privateKey, &privateKey.PublicKey)

		filePEMPrivateKey, err := os.Create(path.Join(pathToPEMs, "pk.pem"))
		if err != nil {
			println("PEM file for the private key could not be generated due to:", err)
			os.Exit(1)
		}
		defer filePEMPrivateKey.Close()

		filePEMPrivateKey.WriteString(pemEncodedPrivateKey)

		filePEMPublicKey, err := os.Create(path.Join(pathToPEMs, "pub.pem"))
		if err != nil {
			println("PEM file for the public key could not be generated due to:", err)
			os.Exit(1)
		}
		defer filePEMPublicKey.Close()

		filePEMPublicKey.WriteString(pemEncodedPublicKey)

		return pemEncodedPrivateKey, pemEncodedPublicKey
	}

	privateKeyBytes, err := os.ReadFile(path.Join(pathToPEMs, "pk.pem"))
	if err != nil {
		println("Private Key could not be read due to:", err)
		os.Exit(1)
	}

	publicKeyBytes, err := os.ReadFile(path.Join(pathToPEMs, "pub.pem"))
	if err != nil {
		println("Public Key could not be read due to:", err)
		os.Exit(1)
	}

	return string(privateKeyBytes), string(publicKeyBytes)
}
