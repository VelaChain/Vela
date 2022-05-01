package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PoolKeyPrefix is the prefix to retrieve all Pool
	PoolKeyPrefix = "Pool/value/"
)

// PoolKey returns the store key to retrieve a Pool from the index fields
func PoolKey(
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
