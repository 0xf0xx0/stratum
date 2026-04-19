package stratum

import "encoding/binary"

// A stratum session id is a uint32 used for the client id, and typically the extranonce1 as well.
type ID uint32

// String returns the [ID] as a big-endian hex string.
func (id ID) String() string {
	return encodeBigEndian(uint32(id))
}

// Bytes returns the [ID] as a 4-byte big-endian array.
func (id ID) Bytes() []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b
}

// DecodeID decodes a big-endian hex string into an [ID].
func DecodeID(s string) (ID, error) {
	x, err := decodeBigEndian(s)
	if err != nil {
		return 0, err
	}

	return ID(x), nil
}
