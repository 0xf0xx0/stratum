package stratum

import "errors"

type ClientShowMessageParams struct {
	Message string
}

func (p *ClientShowMessageParams) FromNotification(n *Notification) error {
	if n.Method != MethodClientShowMessage.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter len; not 1")
	}

	msg, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid message (not string)")
	}
	p.Message = msg
	return nil
}

func (p *ClientShowMessageParams) ToNotification() *Notification {
	return NewNotification(MethodClientShowMessage, []interface{}{p.Message})
}
