package api

import (
	"crypto"
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

var ECDSA = jwt.SigningMethodECDSA{
	Name:      "ECDSASigningMethod",
	Hash:      crypto.SHA256,
	KeySize:   32,
	CurveBits: 256,
}

type Handler struct {
	logger     hclog.Logger
	privateKey *ecdsa.PrivateKey
}
