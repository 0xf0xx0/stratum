package stratum

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// common util funcs

// encodeBigEndian encodes a uint32 as a big-endian hex string.
func encodeBigEndian(n uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)
	return hex.EncodeToString(b)
}

// decodeBigEndian decodes a big-endian hex string into a uint32.
func decodeBigEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("invalid length")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.BigEndian, &x)
	return x, nil
}

// encodeLittleEndian encodes a uint32 as a little-endian hex string.
func encodeLittleEndian(n uint32) string {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, n)
	return hex.EncodeToString(b)
}

// decodeLittleEndian decodes a little-endian hex string into a uint32.
func decodeLittleEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("invalid length")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &x)
	return x, nil
}

// SwapWordEndianness swaps the endianness of each 4-byte word in the buffer.
// Ported from public-pool.
func SwapWordEndianness(buf []byte) []byte {
	swapped := make([]byte, len(buf))

	for i := 0; i < len(buf); i += 4 {
		swapped[i] = buf[i+3]
		swapped[i+1] = buf[i+2]
		swapped[i+2] = buf[i+1]
		swapped[i+3] = buf[i]
	}
	return swapped
}
