package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToHex(t *testing.T) {
	token := "ABCD"

	// Test with tokenFormat set to false
	expectedOutput := "41424344"
	output := StringToHex(token, false)
	assert.Equal(t, expectedOutput, output)

	// Test with tokenFormat set to true
	expectedOutput = "41424344" + "00000000000000000000000000000000"
	output = StringToHex(token, true)
	assert.Equal(t, expectedOutput, output)
}

func TestHexToString(t *testing.T) {
	// test decoding a hex with the Nonstandard Currency Codes
	hex := "41424344" + "00000000000000000000000000000000"

	expectedOutput := "ABCD"
	output, _ := HexToString(hex)
	assert.Equal(t, expectedOutput, output)

	// test decoding a hex without the Nonstandard Currency Codes
	hex = "41424344"

	expectedOutput = "ABCD"
	output, _ = HexToString(hex)
	assert.Equal(t, expectedOutput, output)
}
