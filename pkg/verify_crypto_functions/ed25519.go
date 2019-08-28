package verify_crypto_functions

import "golang.org/x/crypto/ed25519"

func ED25519Sign(privateKey, message []byte) []byte {
	return ed25519.Sign(privateKey, message)
}

func ED25519Verify(publicKey, message, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
