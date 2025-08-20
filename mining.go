package stratum

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
)

type WorkerName string

// Worker represents a miner who is doing work for the pool.
// This would be used in an implementation of a Stratum server and is
// not part of the Stratum protocol.
type Worker struct {
	Name            WorkerName
	SessionID       ID
	ExtraNonce2Size uint32
	VersionMask     *uint32
}

// A share is the data returned by the worker in mining.submit.
type Share struct {
	Name  WorkerName
	JobID string
	MinerShare
}

func (p *Share) Read(r *Request) error {
	if len(r.Params) < 5 || len(r.Params) > 6 {
		return errors.New("invalid format param len")
	}

	name, ok := r.Params[0].(string)
	if !ok {
		return errors.New("invalid format param[0]")
	}

	p.JobID, ok = r.Params[1].(string)
	if !ok {
		return errors.New("invalid format param[1]")
	}

	extraNonce2, ok := r.Params[2].(string)
	if !ok {
		return errors.New("invalid format param[2]")
	}

	time, ok := r.Params[3].(string)
	if !ok {
		return errors.New("invalid format param[3]")
	}

	nonce, ok := r.Params[4].(string)
	if !ok {
		return errors.New("invalid format param[4]")
	}

	if len(r.Params) == 6 {
		rawVersionMask, ok := r.Params[5].(string)
		if !ok {
			return errors.New("invalid format param[5]")
		}


		y, err := hex.DecodeString(rawVersionMask)
		if err != nil {
			return err
		}
		/// this seems to be parsed the wrong way round? i dont know why
		swappedVersionMask := hex.EncodeToString(SwapWordEndianness(y))
		versionMask, err := decodeLittleEndian(swappedVersionMask)
		if err != nil {
			return err
		}

		p.VersionMask = versionMask
	}

	var err error

	p.Nonce, err = decodeBigEndian(nonce)
	if err != nil {
		return err
	}

	p.Time, err = decodeBigEndian(time)
	if err != nil {
		return err
	}

	p.ExtraNonce2, err = hex.DecodeString(extraNonce2)
	if err != nil {
		return err
	}

	p.Name = WorkerName(name)
	return nil
}

// A share is the data returned by the worker. Job + Share = Proof
type MinerShare struct {
	Time        uint32 // proof timestamp
	Nonce       uint32 // gets put into the block header
	ExtraNonce2 []byte // gets put into the coinbase
	VersionMask uint32 // block version + VersionMask = proof version
}

func MakeShare(time uint32, nonce uint32, extraNonce2 uint64) MinerShare {
	n2 := make([]byte, 8)
	binary.BigEndian.PutUint64(n2, extraNonce2)
	return MinerShare{
		Time:        time,
		Nonce:       nonce,
		ExtraNonce2: n2,
	}
}

func MakeShareASICBoost(time uint32, nonce uint32, extraNonce2 uint64, versionMask uint32) MinerShare {
	n2 := make([]byte, 8)
	binary.BigEndian.PutUint64(n2, extraNonce2)
	return MinerShare{
		Time:        time,
		Nonce:       nonce,
		ExtraNonce2: n2,
		VersionMask: versionMask,
	}
}
