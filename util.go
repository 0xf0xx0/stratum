package stratum

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

// common util funcs

func encodeBigEndian(n uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(n))
	return hex.EncodeToString(b)
}

func decodeBigEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, fmt.Errorf("invalid format: %s", s)
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.BigEndian, &x)
	return x, nil
}

func encodeLittleEndian(n uint32) string {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(n))
	return hex.EncodeToString(b)
}

func decodeLittleEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("invalid format")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &x)
	return x, nil
}
