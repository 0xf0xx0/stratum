package stratum

import (
	"errors"
)

type MiningSuggestDifficultyParams struct {
	Difficulty float64
}

func (p *MiningSuggestDifficultyParams) FromRequest(n *Request) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}

func (p *MiningSuggestDifficultyParams) ToRequest(id MessageID) *Request {
	return NewRequest(id, MethodMiningSuggestDifficulty, []interface{}{p.Difficulty})
}
