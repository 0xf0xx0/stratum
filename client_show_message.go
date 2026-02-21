package stratum

import "errors"

type ShowMessageParams struct {
	Message string
}

func (p *ShowMessageParams) Read(n *Notification) error {
	if len(n.Params) != 1 {
		return errors.New("invalid param len (not 1)")
	}

	msg, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid message (not string)")
	}
	p.Message = msg
	return nil
}

func ShowMessage(n ShowMessageParams) *Notification {
	return NewNotification(ClientShowMessage, []interface{}{n.Message})
}
