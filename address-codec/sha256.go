package addresscodec

import (
	"crypto/sha256"

	//nolint:staticcheck // G406: Use of deprecated weak cryptographic primitive (gosec)
	"golang.org/x/crypto/ripemd160"
)

// Returns byte slice of a double hashed given byte slice.
// The given byte slice is SHA256 hashed, then the result is RIPEMD160 hashed.
func Sha256RipeMD160(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	// TODO: Check if this is still needed
	//nolint:gosec // G406: Use of deprecated weak cryptographic primitive (gosec)
	ripemd160 := ripemd160.New()
	ripemd160.Write(sha256.Sum(nil))

	return ripemd160.Sum(nil)
}
