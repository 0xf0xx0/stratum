package stratum

import "errors"

// MiningExtranonceSubscribeParams is a parameter-less message sent from the client to the pool.
// The client should tell the pool it supports [SubscribeExtranonceExtension], then upon successful subscription to the pool send this method.
//
// See: https://github.com/nicehash/Specifications/blob/master/NiceHash_extranonce_subscribe_extension.txt
type MiningExtranonceSubscribeParams struct{}

// FromRequest parses the [MiningExtranonceSubscribeParams] from a [Request].
func (p *MiningExtranonceSubscribeParams) FromRequest(n *Request) error {
	if n.Method != MethodMiningExtranonceSubscribe.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 0 {
		return errors.New("incorrect parameter length; must be 0")
	}
	return nil
}

// ToRequest creates a [Request] from the [MiningExtranonceSubscribeParams].
func (p *MiningExtranonceSubscribeParams) ToRequest(id MessageID) *Request {
	return NewRequest(id, MethodMiningExtranonceSubscribe, []interface{}{})
}
