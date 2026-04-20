package stratum

import (
	"errors"
)

// MiningSetDifficultyParams is sent from the pool to the client.
// It is used to signal the miner to start submitting shares under the new difficulty.
type MiningSetDifficultyParams struct {
	Difficulty float64
}

// FromNotification parses the [MiningSetDifficultyParams] from a [Notification].
func (p *MiningSetDifficultyParams) FromNotification(n *Notification) error {
	if n.Method != MethodMiningSetDifficulty.String() {
		return errors.New("incorrect method")
	}
	if len(n.Params) != 1 {
		return errors.New("incorrect parameter length; must be 2")
	}

	if !validDifficulty(n.Params[0]) {
		return errors.New("invalid difficulty")
	}

	p.Difficulty = n.Params[0].(float64)

	return nil
}

// ToNotification creates a [Notification] from the [MiningSetDifficultyParams].
func (p *MiningSetDifficultyParams) ToNotification() *Notification {
	return NewNotification(MethodMiningSetDifficulty, []interface{}{p.Difficulty})
}
