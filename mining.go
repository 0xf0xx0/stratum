package stratum

import (
	"encoding/hex"
	"errors"
)

// A Share is the data returned by the worker in a mining.submit. Job + Share = Proof
type Share struct {
	Name        string // worker name, like `bc1qfakeaddr.bitaxe`
	JobID       string // Stratum Job ID, should match a mining.notify
	Time        uint32 // proof timestamp
	Nonce       uint32 // gets put into the block header
	ExtraNonce2 []byte // gets put into the coinbase
	VersionMask uint32 // block version + VersionMask = proof version
}

// Read a Share from a Request.
func (p *Share) Read(r *Request) error {
	if len(r.Params) < 5 || len(r.Params) > 6 {
		return errors.New("invalid format param len")
	}

	ok := false

	p.Name, ok = r.Params[0].(string)
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
		return errors.New("param[4]")
	}

	if len(r.Params) == 6 {
		rawVersionMask, ok := r.Params[5].(string)
		if !ok {
			return errors.New("param[5] is not string")
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
	return nil
}
