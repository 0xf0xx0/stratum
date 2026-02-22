package stratum

import (
	"errors"
)

type MiningSuggestDifficultyParams struct {
	Difficulty float64
}

func (p *MiningSuggestDifficultyParams) Read(n *Request) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}

func SuggestDifficultyRequest(id MessageID, r MiningSuggestDifficultyParams) *Request {
	return NewRequest(id, MethodMiningSuggestDifficulty, []interface{}{r.Difficulty})
}
