package stratum

import (
	"errors"
)

type SuggestDifficultyParams struct {
	Difficulty float64
}

func (p *SuggestDifficultyParams) Read(n *Request) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}

func SuggestDifficultyRequest(id MessageID, r SuggestDifficultyParams) Request {
	return NewRequest(id, MiningSuggestDifficulty, []interface{}{r.Difficulty})
}
