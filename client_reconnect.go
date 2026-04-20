package stratum

import "errors"

// ClientReconnectParams is sent from the pool to the client, and tells the client when and where to reconnect.
//
// If client.reconnect is sent without parameters, the miner is to assume it's to reconnect to the same port and URL.
type ClientReconnectParams struct {
	Hostname       string
	Port, Waittime uint16
}

// FromNotification parses the [ClientReconnectParams] from a [Notification].
func (p *ClientReconnectParams) FromNotification(n *Notification) error {
	if n.Method != MethodClientReconnect.String() {
		return errors.New("incorrect method")
	}
	/// "If client.reconnect is sent without parameters, the miner is to assume it's to reconnect to the same port and URL."
	/// - ck, https://bitcointalk.org/index.php?topic=557866.msg6989610#msg6989610
	if len(n.Params) == 0 {
		return nil
	}
	if len(n.Params) < 2 || len(n.Params) > 3 {
		return errors.New("incorrect parameter len; not 2 or 3")
	}

	hostname, ok := n.Params[0].(string)
	if !ok {
		return errors.New("invalid hostname (not string)")
	}
	p.Hostname = hostname

	port, ok := n.Params[1].(float64)
	if !ok {
		return errors.New("invalid port (not uint16)")
	}
	p.Port = uint16(port)

	if len(n.Params) == 3 {
		waittime, ok := n.Params[2].(float64)
		if !ok {
			return errors.New("invalid waittime (not uint16)")
		}
		p.Waittime = uint16(waittime)
	}
	return nil
}

// ToNotification creates a [Notification] from the [ClientReconnectParams].
func (p *ClientReconnectParams) ToNotification() *Notification {
	params := make([]any, 3)
	params[0] = p.Hostname
	params[1] = p.Port

	if p.Waittime > 0 {
		params[2] = p.Waittime
	}
	return NewNotification(MethodClientReconnect, params)
}
