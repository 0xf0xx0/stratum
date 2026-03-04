package stratum

import "errors"

type MiningExtranonceSubscribeParams struct{}

func (p *MiningExtranonceSubscribeParams) FromRequest(n *Request) error {
	if n.Method != MethodMiningExtranonceSubscribe.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 0 {
		return errors.New("incorrect parameter length, XNSUB has 0 params")
	}
	return nil
}

func (p *MiningExtranonceSubscribeParams) ToRequest(id MessageID) *Request {
	return NewRequest(id, MethodMiningExtranonceSubscribe, []interface{}{})
}
