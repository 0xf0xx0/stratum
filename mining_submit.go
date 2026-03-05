package stratum

import (
	"encoding/hex"
	"errors"
)

// Alias for [MiningSubmitParams].
type Share = MiningSubmitParams
type SubmitResult BooleanResult

// A MiningSubmitParams is the data returned by the worker in a mining.submit. Job + MiningSubmitParams = Proof
type MiningSubmitParams struct {
	Name        string // worker name, like `bc1qfakeaddr.bitaxe`
	JobID       string // Stratum Job ID, must match a mining.notify
	Time        uint32 // proof timestamp
	Nonce       uint32 // gets put into the block header
	ExtraNonce2 []byte // gets put into the coinbase
	VersionMask uint32 // block version + VersionMask = proof version
}

// FromRequest creates a MiningSubmitParams from a Request.
func (p *MiningSubmitParams) FromRequest(r *Request) error {
	if r.Method != MethodMiningSubmit.String() {
		return errors.New("incorrect method")
	}
	if len(r.Params) < 5 || len(r.Params) > 6 {
		return errors.New("incorrect parameter length; must be 5 or 6")
	}

	ok := false

	p.Name, ok = r.Params[0].(string)
	if !ok {
		return errors.New("invalid name (not string)")
	}

	p.JobID, ok = r.Params[1].(string)
	if !ok {
		return errors.New("invalid jobid (not string)")
	}

	extraNonce2, ok := r.Params[2].(string)
	if !ok {
		return errors.New("invalid extranonce2 (not string)")
	}

	time, ok := r.Params[3].(string)
	if !ok {
		return errors.New("invalid time (not string)")
	}

	nonce, ok := r.Params[4].(string)
	if !ok {
		return errors.New("invalid nonce (not string)")
	}

	if len(r.Params) == 6 {
		rawVersionMask, ok := r.Params[5].(string)
		if !ok {
			return errors.New("invalid version_mask (not string)")
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
func (p *MiningSubmitParams) ToRequest(id MessageID) *Request {
	var sx []interface{}
	if p.VersionMask != 0 {
		sx = make([]interface{}, 6)
		sx[5] = encodeLittleEndian(p.VersionMask)
	} else {
		sx = make([]interface{}, 5)
	}

	sx[0] = string(p.Name)
	sx[1] = p.JobID
	sx[2] = hex.EncodeToString(p.ExtraNonce2)
	sx[3] = encodeBigEndian(p.Time)
	sx[4] = encodeBigEndian(p.Nonce)

	return NewRequest(id, MethodMiningSubmit, sx)
}
