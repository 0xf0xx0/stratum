package stratum

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced.
// we also use ID for job ids. [kbnchk:] bad idea jobID is not always hex-encoded uint32, removed
type ID uint32
// calls EncodeID()
func (id ID) String() string {
	return EncodeID(id)
}

func EncodeID(id ID) string {
	return encodeBigEndian(uint32(id))
}
func DecodeID(s string) (ID, error) {
	x, err := decodeBigEndian(s)
	if err != nil {
		return 0, err
	}

	return ID(x), nil
}
