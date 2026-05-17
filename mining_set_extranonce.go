package stratum

import "errors"

// MiningSetExtranonceParams is sent from the pool to the client.
// It is used to change the clients extranonce1, and optionally extranonce2 size, on the fly.
type MiningSetExtranonceParams struct {
	ExtraNonce1     ID
	ExtraNonce2Size int
}

// FromNotification parses the [MiningSetExtranonceParams] from a [Notification].
func (p *MiningSetExtranonceParams) FromNotification(n *Notification) error {
	if n.Method != MethodMiningSetExtranonce.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 2 {
		return errors.New("incorrect parameter length; must be 2")
	}
	rawExtranonce1, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid extranonce1 (not string)")
	}
	decodedExtranonce1, err := DecodeID(rawExtranonce1)
	if err != nil {
		return err
	}
	p.ExtraNonce1 = decodedExtranonce1

	en2size, ok := n.Params[1].(float64)
	if !ok {
		return errors.New("couldnt cast extranonce2size to float64")
	}
	p.ExtraNonce2Size = int(en2size)

	return nil
}

// ToNotification creates a [Notification] from the [MiningSetExtranonceParams].
func (p *MiningSetExtranonceParams) ToNotification(id MessageID) *Notification {
	return NewNotification(MethodMiningSetExtranonce, []any{p.ExtraNonce1.String(), p.ExtraNonce2Size})
}
