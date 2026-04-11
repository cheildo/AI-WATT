package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// HMACSign computes an HMAC-SHA256 signature of payload using key.
// Used by the Veriflow agent and backend to verify telemetry authenticity.
func HMACSign(key, payload []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

// HMACVerify returns true if the provided signature matches the expected HMAC.
func HMACVerify(key, payload []byte, signature string) bool {
	expected := HMACSign(key, payload)
	return hmac.Equal([]byte(expected), []byte(signature))
}

// Keccak256 computes the keccak256 hash of data (compatible with Solidity keccak256).
func Keccak256(data []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	return h.Sum(nil)
}

// Keccak256Hex returns the keccak256 hash of data as a hex string.
func Keccak256Hex(data []byte) string {
	return hex.EncodeToString(Keccak256(data))
}
