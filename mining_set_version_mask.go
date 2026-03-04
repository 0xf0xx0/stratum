package stratum

import "errors"

type MiningSetVersionMaskParams struct {
	Mask uint32
}

func (p *MiningSetVersionMaskParams) FromNotification(n *Notification) error {
	if n.Method != MethodMiningSetVersionMask.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 1 {
		return errors.New("invalid format")
	}

	mask, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid format")
	}

	var err error
	p.Mask, err = decodeLittleEndian(mask)
	if err != nil {
		return err
	}

	return nil
}

func SetVersionMask(u uint32) *Notification {
	return NewNotification(MethodMiningSetVersionMask, []interface{}{encodeLittleEndian(u)})
}
