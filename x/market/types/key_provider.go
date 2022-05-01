package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProviderKeyPrefix is the prefix to retrieve all Provider
	ProviderKeyPrefix = "Provider/value/"
)

// ProviderKey returns the store key to retrieve a Provider from the index fields
func ProviderKey(
	denomA string,
	denomB string,
) []byte {
	var key []byte

	denomABytes := []byte(denomA)
	key = append(key, denomABytes...)
	key = append(key, []byte("/")...)

	denomBBytes := []byte(denomB)
	key = append(key, denomBBytes...)
	key = append(key, []byte("/")...)

	return key
}
