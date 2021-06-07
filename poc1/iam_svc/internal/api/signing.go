package api

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func generateECDSAKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privKey, &privKey.PublicKey
}
