package stratum

// Difficulty must be given as a float64.
// TODO: fix params
func validDifficulty(u interface{}) bool {
	switch d := u.(type) {
	case float64:
		return d > 0
	default:
		return false
	}
}
