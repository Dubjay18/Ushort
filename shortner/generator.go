// Package shortener provides functionality to generate short URLs.
package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// Generator struct doesn't hold any state and is used to implement the ShortLinkGenerator interface.
type Generator struct {
}

// ShortLinkGenerator interface defines the method that our URL shortener must implement.
type ShortLinkGenerator interface {
	GenerateShortLink(initialLink string) string
}

// NewGenerator function returns a new instance of Generator.
func NewGenerator() *Generator {
	return &Generator{}
}

// GenerateShortLink method generates a short URL from the given initial URL.
func (g *Generator) GenerateShortLink(initialLink string) string {
	// Generate a SHA-256 hash of the initial URL.
	urlHashBytes := sha256Of(initialLink + "jay")
	// Convert the hash to a big integer and then to a uint64.
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	// Base58 encode the uint64 to get a string.
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	// Return the first 8 characters of the string as the short URL.
	return finalString[:8]
}

// sha256Of function generates a SHA-256 hash of the given input string.
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base58Encoded function encodes the given bytes into a Base58 string.
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
