package stratum

import (
	"errors"
)

type MiningSetDifficultyParams struct {
	Difficulty float64
}

func (p *MiningSetDifficultyParams) FromNotification(n *Notification) error {
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}
func (p *MiningSetDifficultyParams) ToNotification() *Notification {
	return NewNotification(MethodMiningSetDifficulty, []interface{}{p.Difficulty})
}

/// TODO: figure out how to remove this
func SetDifficulty(d float64) *Notification {
	return NewNotification(MethodMiningSetDifficulty, []interface{}{d})
}
