package Stratum

import (
	"encoding/hex"
	"errors"
	"reflect"
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
		return errors.New("invalid format")
	}

	var ok bool
	p.JobID, ok = n.Params[0].(string)
	if !ok {
		return errors.New("invalid format")
	}

	digest, ok := n.Params[1].(string)
	if !ok {
		return errors.New("invalid format")
	}

	gtx1, ok := n.Params[2].(string)
	if !ok {
		return errors.New("invalid format")
	}

	gtx2, ok := n.Params[3].(string)
	if !ok {
		return errors.New("invalid format")
	}

	/// FIXME
	path := make([]string, 0, 16)
	rv := reflect.ValueOf(n.Params[4])
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			path = append(path, rv.Index(i).Elem().String())
		}
	} else {
		return errors.New("invalid format")
	}

	// ok, path := n.Params[4].([]string)
	// if !ok {
	// 	return errors.New("invalid format")
	// }

	version, ok := n.Params[5].(string)
	if !ok {
		return errors.New("invalid format")
	}

	bits, ok := n.Params[6].(string)
	if !ok {
		return errors.New("invalid format")
	}

	timestamp, ok := n.Params[7].(string)
	if !ok {
		return errors.New("invalid format")
	}
	ts, err := decodeBigEndian(timestamp)
	if err != nil {
		return errors.New("invalid format")
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
	if err != nil || len(p.Bits) != 4 {
		return errors.New("invalid format")
	}

	p.CoinbasePart1, err = hex.DecodeString(gtx1)
	if err != nil {
		return errors.New("invalid format")
	}

	p.CoinbasePart2, err = hex.DecodeString(gtx2)
	if err != nil {
		return errors.New("invalid format")
	}

	p.Version, err = decodeBigEndian(version)
	if err != nil {
		return errors.New("invalid format")
	}


	p.MerkleBranches = make([][]byte, len(path))
	for i := 0; i < len(path); i++ {
		p.MerkleBranches[i], err = hex.DecodeString(path[i])
		if err != nil || len(p.PrevBlockHash) != 32 {
			return errors.New("invalid format")
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
