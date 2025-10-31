package stratum

import "encoding/binary"

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced as extranonce1.
type ID uint32

// returns the id as a big-endian hex string
func (id ID) String() string {
	return encodeBigEndian(uint32(id))
}

// 4 byte array, big-endian
func (id ID) Bytes() []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b
}

// decodes an id from a big-endian hex string
func DecodeID(s string) (ID, error) {
	x, err := decodeBigEndian(s)
	if err != nil {
		return 0, err
	}

	return ID(x), nil
}
