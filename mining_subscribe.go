package stratum

import (
	"errors"
	"math"
)

type MiningSubscribeParams struct {
	UserAgent   string // required
	ExtraNonce1 *ID // optional extranonce subscription
}

func (p *MiningSubscribeParams) Read(r *Request) error {
	l := len(r.Params)
	if l == 0 || l > 2 {
		return errors.New("invalid parameter length; must be 1 or 2")
	}

	var ok bool
	p.UserAgent, ok = r.Params[0].(string)
	if !ok {
		return errors.New("invalid user agent format")
	}

	if l == 1 {
		p.ExtraNonce1 = nil
		return nil
	}

	idstr, ok := r.Params[1].(string)
	if !ok {
		return errors.New("invalid session id format")
	}

	id, err := DecodeID(idstr)
	if err != nil {
		return err
	}

	p.ExtraNonce1 = &id
	return nil
}

func SubscribeRequest(id MessageID, r MiningSubscribeParams) *Request {
	if r.ExtraNonce1 == nil {
		return NewRequest(id, MethodMiningSubscribe, []interface{}{r.UserAgent})
	}

	return NewRequest(id, MethodMiningSubscribe, []interface{}{r.UserAgent, r.ExtraNonce1.String()})
}

// A Subscription is a 2-element json array containing a method
// and a session id.
type Subscription struct {
	Method    Method
	SessionID ID
}

type SubscribeResult struct {
	Subscriptions   []Subscription
	ExtraNonce1     ID
	ExtraNonce2Size uint32
}

func (p *SubscribeResult) Read(r *Response) error {
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
	p.Subscriptions = make([]Subscription, len(subscriptions))
	for i := 0; i < len(subscriptions); i++ {
		sub := subscriptions[i].([]interface{})
		if len(sub) != 2 {
			return errors.New("invalid subscriptions format")
		}

		p.Subscriptions[i].Method, err = DecodeMethod(sub[0].(string))
		if err != nil {
			return err
		}

		p.Subscriptions[i].SessionID, err = DecodeID(sub[1].(string))
		if err != nil {
			return err
		}
	}

	p.ExtraNonce1, err = DecodeID(idstr)
	if err != nil {
		return errors.New("invalid session id")
	}

	return nil
}

func SubscribeResponse(m MessageID, r SubscribeResult) *Response {
	subscriptions := make([][]string, len(r.Subscriptions))
	for i := 0; i < len(r.Subscriptions); i++ {
		subscriptions[i] = make([]string, 2)

		method, err := EncodeMethod(r.Subscriptions[i].Method)
		if err != nil {
			/// TODO: return error? i dont wanna change just this function sig
			return NewResponse(0, nil)
		}

		subscriptions[i][0] = method
		subscriptions[i][1] = r.Subscriptions[i].SessionID.String()
	}

	result := make([]interface{}, 3)
	result[0] = subscriptions
	result[1] = r.ExtraNonce1.String()
	result[2] = r.ExtraNonce2Size

	return NewResponse(m, result)
}
