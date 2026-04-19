package stratum

import (
	"errors"
	"math"
)

type MiningSubscribeParams struct {
	UserAgent   string // required
	ExtraNonce1 *ID    // optional extranonce subscription
}

// FromRequest parses the [MiningSubscribeParams] from a [Request].
func (p *MiningSubscribeParams) FromRequest(r *Request) error {
	if r.Method != MethodMiningSubscribe.String() {
		return errors.New("incorrect method")
	}
	l := len(r.Params)
	if l == 0 || l > 2 {
		return errors.New("invalid parameter length; must be 1 or 2")
	}

	var ok bool
	p.UserAgent, ok = r.Params[0].(string)
	if !ok {
		return errors.New("invalid user agent (not string)")
	}

	if l == 1 {
		p.ExtraNonce1 = nil
		return nil
	}

	idstr, ok := r.Params[1].(string)
	if !ok {
		return errors.New("invalid session id (string)")
	}

	id, err := DecodeID(idstr)
	if err != nil {
		return err
	}

	p.ExtraNonce1 = &id
	return nil
}

// ToRequest creates a [Request] from the [MiningSubscribeParams].
func (p *MiningSubscribeParams) ToRequest(id MessageID) *Request {
	if p.ExtraNonce1 == nil {
		return NewRequest(id, MethodMiningSubscribe, []interface{}{p.UserAgent})
	}
	return NewRequest(id, MethodMiningSubscribe, []interface{}{p.UserAgent, p.ExtraNonce1.String()})
}

// A MiningSubscription is a 2-element json array containing a method and a session id.
type MiningSubscription struct {
	Method    Method
	SessionID ID
}

type MiningSubscribeResult struct {
	Subscriptions   []MiningSubscription
	ExtraNonce1     ID
	ExtraNonce2Size uint32
}

// FromResponse parses the [MiningSubscribeResult] from a [Response].
func (p *MiningSubscribeResult) FromResponse(r *Response) error {
	result, ok := r.Result.([]interface{})
	if !ok {
		return errors.New("invalid result type; should be array")
	}

	if len(result) != 3 {
		return errors.New("invalid parameter length; must be 3")
	}

	subscriptions := result[0].([]interface{})

	idstr, ok := result[1].(string)
	if !ok {
		return errors.New("invalid session id")
	}

	extraNonce2Size := uint64(result[2].(float64))

	if extraNonce2Size > math.MaxUint32 {
		return errors.New("extraNonce2_size too big")
	}

	p.ExtraNonce2Size = uint32(extraNonce2Size)

	var err error
	p.Subscriptions = make([]MiningSubscription, len(subscriptions))
	for i, s := range subscriptions {
		sub := s.([]interface{})
		if len(sub) != 2 {
			return errors.New("incorrect subscription length; must be 2")
		}

		p.Subscriptions[i].Method = DecodeMethod(sub[0].(string))

		p.Subscriptions[i].SessionID, err = DecodeID(sub[1].(string))
		if err != nil {
			return err
		}
	}

	p.ExtraNonce1, err = DecodeID(idstr)
	if err != nil {
		return err
	}

	return nil
}

// ToResponse creates a [Response] from the [MiningSubscribeResult].
func (p *MiningSubscribeResult) ToResponse(m MessageID) *Response {
	subscriptions := make([][]string, len(p.Subscriptions))
	for i := 0; i < len(p.Subscriptions); i++ {
		subscriptions[i] = make([]string, 2)

		method, err := EncodeMethod(p.Subscriptions[i].Method)
		if err != nil {
			/// TODO: return error? i dont wanna change just this function sig
			return NewResponse(0, nil)
		}

		subscriptions[i][0] = method
		subscriptions[i][1] = p.Subscriptions[i].SessionID.String()
	}

	result := make([]interface{}, 3)
	result[0] = subscriptions
	result[1] = p.ExtraNonce1.String()
	result[2] = p.ExtraNonce2Size

	return NewResponse(m, result)
}
