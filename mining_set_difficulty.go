package stratum

import (
	"errors"
)

type SetDifficultyParams struct {
	Difficulty float64
}

func (p *SetDifficultyParams) Read(n *Notification) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}

func SetDifficulty(d float64) Notification {
	return NewNotification(MiningSetDifficulty, []interface{}{d})
}
