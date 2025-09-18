package stratum

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type NotifyParams struct {
	JobID          string
	PrevBlockHash  *chainhash.Hash
	CoinbasePart1  []byte
	CoinbasePart2  []byte
	MerkleBranches [][]byte
	Version        uint32
	Bits           []byte
	Timestamp      time.Time
	Clean          bool
}

func (p *NotifyParams) Read(n *Notification) error {
	if len(n.Params) != 9 {
		return errors.New("invalid param len (not 9)")
	}

	var ok bool
	p.JobID, ok = n.Params[0].(string)
	if !ok {
		return errors.New("invalid jobid (not string)")
	}

	digest, ok := n.Params[1].(string)
	if !ok {
		return errors.New("invalid prevblockhash (not string)")
	}

	gtx1, ok := n.Params[2].(string)
	if !ok {
		return errors.New("invalid coinbasept1 (not string)")
	}

	gtx2, ok := n.Params[3].(string)
	if !ok {
		return errors.New("invalid coinbasept2 (not string)")
	}

	mb, ok := n.Params[4].([]interface{})
	if !ok {
		return errors.New("invalid merkle branches type (not []string)")
	}
	branches := make([]string, len(mb))
	for i := range mb {
		s, ok := mb[i].(string)
		if !ok {
			return errors.New("invalid merkle branch type (not string)")
		}
		branches[i] = s
	}

	version, ok := n.Params[5].(string)
	if !ok {
		return errors.New("invalid version (not string)")
	}

	bits, ok := n.Params[6].(string)
	if !ok {
		return errors.New("invalid bits (not string)")
	}

	timestamp, ok := n.Params[7].(string)
	if !ok {
		return errors.New("invalid timestamp (not string)")
	}
	ts, err := decodeBigEndian(timestamp)
	if err != nil {
		return errors.New("invalid timestamp: "+err.Error())
	}
	p.Timestamp = time.Unix(int64(ts), 0)

	p.Clean, ok = n.Params[8].(bool)
	if !ok {
		return errors.New("invalid format")
	}

	p.PrevBlockHash, err = chainhash.NewHashFromStr(digest)
	if err != nil || len(p.PrevBlockHash) != 32 {
		return errors.New("invalid format")
	}

	p.Bits, err = hex.DecodeString(bits)
	if err != nil {
		return err
	}
	if len(p.Bits) != 4 {
		return errors.New("invalid bits param size (not 4 bytes)")
	}

	p.CoinbasePart1, err = hex.DecodeString(gtx1)
	if err != nil {
		return errors.New("invalid coinbase pt 1: "+err.Error())
	}

	p.CoinbasePart2, err = hex.DecodeString(gtx2)
	if err != nil {
		return errors.New("invalid coinbase pt 2: "+err.Error())
	}

	p.Version, err = decodeBigEndian(version)
	if err != nil {
		return errors.New("invalid version: "+err.Error())
	}

	p.MerkleBranches = make([][]byte, len(branches))
	for i := 0; i < len(branches); i++ {
		p.MerkleBranches[i], err = hex.DecodeString(branches[i])
		if err != nil || len(p.MerkleBranches[i]) != 32 {
			return errors.New("invalid merkle branch length (not 32)")
		}
	}

	return nil
}

func Notify(n NotifyParams) Notification {
	params := make([]interface{}, 9)

	params[0] = n.JobID
	params[1] = hex.EncodeToString(SwapWordEndianness(n.PrevBlockHash[:]))
	params[2] = hex.EncodeToString(n.CoinbasePart1)
	params[3] = hex.EncodeToString(n.CoinbasePart2)

	path := make([]string, len(n.MerkleBranches))
	for i := 0; i < len(n.MerkleBranches); i++ {
		path[i] = hex.EncodeToString(n.MerkleBranches[i])
	}

	params[4] = path
	params[5] = encodeBigEndian(n.Version)
	params[6] = hex.EncodeToString(n.Bits)
	params[7] = encodeBigEndian(uint32(n.Timestamp.Unix()))
	params[8] = n.Clean

	return NewNotification(MiningNotify, params)
}

// ported from public-pool
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
