package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LiqProvKeyPrefix is the prefix to retrieve all LiqProv
	LiqProvKeyPrefix = "LiqProv/value/"
)

// LiqProvKey returns the store key to retrieve a LiqProv from the index fields
func LiqProvKey(
	poolName string,
	address string,
) []byte {
	var key []byte

	poolNameBytes := []byte(poolName)
	key = append(key, poolNameBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
