package stratum

import (
	"errors"
)

// MiningSuggestDifficultyParams is sent from the client to the pool.
// It is used once to suggest a starting difficulty, and may be ignored by the pool.
// If accepted, the pool will echo the suggested difficulty in a [MethodMiningSetDifficulty].
type MiningSuggestDifficultyParams struct {
	Difficulty float64
}

// FromRequest parses the [MiningSuggestDifficultyParams] from a [Request].
func (p *MiningSuggestDifficultyParams) FromRequest(r *Request) error {
	if r.Method != MethodMiningSuggestDifficulty.String() {
		return errors.New("incorrect method")
	}
	if len(r.Params) != 1 {
		return errors.New("incorrect parameter length; must be 1")
	}

	if !validDifficulty(r.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = r.Params[0].(float64)

	return nil
}

// ToRequest creates a [Request] from the [MiningSuggestDifficultyParams].
func (p *MiningSuggestDifficultyParams) ToRequest(id MessageID) *Request {
	return NewRequest(id, MethodMiningSuggestDifficulty, []any{p.Difficulty})
}
