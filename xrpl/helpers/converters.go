package helpers

import (
	"encoding/hex"
	"strings"
)

// StringToHex converts a string to a hexadecimal representation.
// If isNonStandardCurrency is true, the hexadecimal representation is padded
// with trailing zeros up to a length of 40 characters. This is to support the non-standard currency codes for the XRPL.
// See https://xrpl.org/docs/references/protocol/data-types/currency-formats#nonstandard-currency-codes
func StringToHex(input string, isNonStandardCurrency bool) string {
	// Convert the string to bytes
	bytes := []byte(input)

	// Convert bytes to hexadecimal representation
	hexString := hex.EncodeToString(bytes)

	if isNonStandardCurrency {
		// Pad end the hex string with trailing zeros up to a length of 40 characters
		hexString = padEnd(hexString, 40, "0")
	}

	return hexString
}

// HexToString converts a hexadecimal representation to a string.
func HexToString(input string) (string, error) {
	// Convert the hexadecimal representation to bytes
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	// Remove null bytes from the byte slice
	trimmedBytes := bytes[:bytesIndex(bytes, 0)]

	// Convert bytes to string
	str := string(trimmedBytes)

	return str, nil
}

// bytesIndex returns the index of the first occurrence of the given value in the byte slice.
// If the value is not found, it returns the length of the slice.
func bytesIndex(slice []byte, value byte) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return len(slice)
}

// padEnd pads the string `s` with the character `padChar` on the right until it reaches `length`.
func padEnd(s string, length int, padChar string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(padChar, length-len(s))
	return s + padding
}
