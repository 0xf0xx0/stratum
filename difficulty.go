package stratum

// Internal helper to assert difficulty as a positive, non-zero float64
func validDifficulty(u interface{}) bool {
	switch d := u.(type) {
	case float64:
		return d > 0
	default:
		return false
	}
}
