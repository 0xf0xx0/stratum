package stratum

import "errors"

type ReconnectParams struct {
	Hostname       string
	Port, Waittime uint16
}

// this assumes hostname and port are required, and waittime is optional
func (p *ReconnectParams) Read(n *Notification) error {
	/// "If client.reconnect is sent without parameters, the miner is to assume it's to reconnect to the same port and URL."
	/// - ck, https://bitcointalk.org/index.php?topic=557866.msg6989610#msg6989610
	if len(n.Params) == 0 {
		return nil
	}
	if len(n.Params) < 2 || len(n.Params) > 3 {
		return errors.New("invalid param len (not 2 or 3)")
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

func Reconnect(n ReconnectParams) *Notification {
	params := make([]interface{}, 3)
	params[0] = n.Hostname
	params[1] = n.Port

	if n.Waittime > 0 {
		params[2] = n.Waittime
	}
	return NewNotification(ClientReconnect, params)
}
