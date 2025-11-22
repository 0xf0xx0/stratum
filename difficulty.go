package stratum

// Difficulty can be given as a uint or a float.
func validDifficulty(u interface{}) bool {
	switch d := u.(type) {
	case float64:
		return d > 0
	default:
		return false
	}
}
