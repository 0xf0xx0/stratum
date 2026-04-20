package stratum

import "errors"

// ClientShowMessageParams is sent from the pool to the client, with a message for the client to display to the user.
type ClientShowMessageParams struct {
	Message string
}

// FromNotification parses the [ClientShowMessageParams] from a [Notification].
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

// ToNotification creates a [Notification] from the [ClientShowMessageParams].
func (p *ClientShowMessageParams) ToNotification() *Notification {
	return NewNotification(MethodClientShowMessage, []any{p.Message})
}
