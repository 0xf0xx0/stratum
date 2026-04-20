package stratum

import "errors"
// MiningSetVersionMaskParams is sent from the client to the pool.
// It is used to set the clients version rolling mask, if it supports [VersionRollingExtension].
// This message can be sent any time after successful setup of the version rolling by [MethodMiningConfigure] message.
type MiningSetVersionMaskParams struct {
	Mask uint32
}

// FromNotification parses the [MiningSetVersionMaskParams] from a [Notification].
func (p *MiningSetVersionMaskParams) FromNotification(n *Notification) error {
	if n.Method != MethodMiningSetVersionMask.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length; must be 1")
	}

	mask, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid mask (not string)")
	}

	var err error
	p.Mask, err = decodeLittleEndian(mask)
	if err != nil {
		return err
	}

	return nil
}

// ToNotification creates a [Notification] from the [MiningSetVersionMaskParams].
func (p *MiningSetVersionMaskParams) ToNotification() *Notification {
	return NewNotification(MethodMiningSetVersionMask, []interface{}{encodeLittleEndian(p.Mask)})
}
