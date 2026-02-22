package stratum

import (
	"errors"
	"strings"
)

type MiningAuthorizeParams struct {
	Username string
	// optional, typically appened to the username with `.`
	Worker string
	// Password is optional. Pools don't necessarily require a miner to log in to mine.
	Password string
}

func (p *MiningAuthorizeParams) Read(r *Request) error {
	l := len(r.Params)
	if l == 0 || l > 2 {
		return errors.New("invalid parameter length; must be 1 or 2")
	}

	username, ok := r.Params[0].(string)
	if !ok {
		return errors.New("invalid username format")
	}

	split := strings.Split(username, ".")
	p.Username = split[0]
	if len(split) > 1 {
		p.Worker = split[1]
	}

	if l == 1 {
		p.Password = ""
		return nil
	}

	password, ok := r.Params[1].(string)
	if !ok {
		return errors.New("invalid password format")
	}

	p.Password = password
	return nil
}

func AuthorizeRequest(id MessageID, r MiningAuthorizeParams) *Request {
	username := r.Username
	if r.Worker != "" {
		username += "." + r.Worker
	}
	if r.Password == "" {
		return NewRequest(id, MethodMiningAuthorize, []interface{}{username})
	}

	return NewRequest(id, MethodMiningAuthorize, []interface{}{username, r.Password})
}

type AuthorizeResult BooleanResult

func AuthorizeResponse(id MessageID, b bool) *Response {
	return NewBooleanResponse(id, b)
}
