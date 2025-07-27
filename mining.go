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

		versionMask, err := decodeLittleEndian(rawVersionMask)
		if err != nil {
			return err
		}

		p.VersionMask = &versionMask
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
	Time        uint32
	Nonce       uint32
	ExtraNonce2 []byte
	// TODO: does this need to be a pointer?
	VersionMask *uint32
}

func MakeShare(time uint32, nonce uint32, extraNonce2 uint64) MinerShare {
	n2 := make([]byte, 8)
	binary.BigEndian.PutUint64(n2, extraNonce2)
	return MinerShare{
		Time:        time,
		Nonce:       nonce,
		ExtraNonce2: n2,
		VersionMask: nil}
}

func MakeShareASICBoost(time uint32, nonce uint32, extraNonce2 uint64, versionMask uint32) MinerShare {
	bits := new(uint32)
	*bits = versionMask
	n2 := make([]byte, 8)
	binary.BigEndian.PutUint64(n2, extraNonce2)
	return MinerShare{
		Time:        time,
		Nonce:       nonce,
		ExtraNonce2: n2,
		VersionMask: bits}
}
