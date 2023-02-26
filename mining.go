package Stratum

import (
	"encoding/hex"
	"errors"

	"github.com/DanielKrawisz/go-work"
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
	JobID ID
	work.Share
}

func (p *Share) Read(r *Request) error {
	if len(r.Params) < 5 || len(r.Params) > 6 {
		return errors.New("invalid format")
	}

	name, ok := r.Params[0].(string)
	if !ok {
		return errors.New("invalid format")
	}

	jobID, ok := r.Params[1].(string)
	if !ok {
		return errors.New("invalid format")
	}

	extraNonce2, ok := r.Params[2].(string)
	if !ok {
		return errors.New("invalid format")
	}

	time, ok := r.Params[3].(string)
	if !ok {
		return errors.New("invalid format")
	}

	nonce, ok := r.Params[4].(string)
	if !ok {
		return errors.New("invalid format")
	}

	if len(r.Params) == 6 {
		gpr, ok := r.Params[5].(string)
		if !ok {
			return errors.New("invalid format")
		}

		GPR, err := decodeLittleEndian(gpr)
		if err != nil {
			return err
		}

		p.GeneralPurposeBits = &GPR
	}

	var err error

	p.JobID, err = decodeID(jobID)
	if err != nil {
		return err
	}

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
