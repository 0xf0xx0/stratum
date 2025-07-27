package Stratum

import (
	"errors"
)

type SuggestDifficultyParams struct {
	Difficulty Difficulty
}

func (p *SuggestDifficultyParams) Read(n *Request) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !ValidDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	// TODO: float?
	p.Difficulty = n.Params[0]

	return nil
}

func SuggestDifficultyRequest(id MessageID, r SuggestDifficultyParams) Request {
	return NewRequest(id, MiningSuggestDifficulty, []interface{}{r.Difficulty})
}
