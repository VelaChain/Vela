package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FeeMapKeyPrefix is the prefix to retrieve all FeeMap
	FeeMapKeyPrefix = "FeeMap/value/"
)

// FeeMapKey returns the store key to retrieve a FeeMap from the index fields
func FeeMapKey(
	poolName string,
) []byte {
	var key []byte

	poolNameBytes := []byte(poolName)
	key = append(key, poolNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
