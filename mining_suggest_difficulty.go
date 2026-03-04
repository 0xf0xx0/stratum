package stratum

import (
	"errors"
)

type MiningSuggestDifficultyParams struct {
	Difficulty float64
}

func (p *MiningSuggestDifficultyParams) FromRequest(r *Request) error {
	if r.Method != MethodMiningSuggestDifficulty.String() {
		return errors.New("incorrect method")
	}
	if len(r.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(r.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = r.Params[0].(float64)

	return nil
}

func (p *MiningSuggestDifficultyParams) ToRequest(id MessageID) *Request {
	return NewRequest(id, MethodMiningSuggestDifficulty, []interface{}{p.Difficulty})
}
