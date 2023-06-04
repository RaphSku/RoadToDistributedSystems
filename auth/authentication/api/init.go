package api

import (
	"crypto/ecdsa"

	"github.com/hashicorp/go-hclog"
)

func (h *Handler) Initialize(logger hclog.Logger, privateKey *ecdsa.PrivateKey) error {
	h.logger = logger
	h.privateKey = privateKey

	return nil
}
