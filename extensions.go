package stratum

import (
	"errors"
)

// https://web.archive.org/web/20260213021043/https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki
type Extension uint8

const (
	// UnknownExtension is unknown.
	UnknownExtension Extension = iota
	// VersionRollingExtension allows the miner to change the value of some bits in the version field in the block header.
	// BIP-320 defines the standard version rolling mask as `0x1fffe000`.
	// The pool will AND the miner-provided mask with the standard and return the result.
	VersionRollingExtension
	// MinimumDifficultyExtension allows miner to request a minimum difficulty for the connected machine.
	// It solves a problem in the original stratum protocol where there is no way how to communicate hard limit of the connected device.
	MinimumDifficultyExtension
	// SubscribeExtranonceExtension is a parameter-less extension.
	// Miner advertises its capability of receiving message "mining.set_extranonce" message (useful for hash rate routing scenarios).
	SubscribeExtranonceExtension
	// InfoExtension allows the miner to provide additional text-based information about itself.
	InfoExtension
)

// EncodeExtension converts an [Extension] to its string name.
func EncodeExtension(m Extension) (string, error) {
	switch m {
	case VersionRollingExtension:
		return "version-rolling", nil
	case MinimumDifficultyExtension:
		return "minimum-difficulty", nil
	case SubscribeExtranonceExtension:
		return "subscribe-extranonce", nil
	case InfoExtension:
		return "info", nil
	default:
		return "", errors.New("unknown stratum extension")
	}
}

// DecodeExtension converts a string name to an [Extension].
func DecodeExtension(m string) Extension {
	switch m {
	case "version-rolling":
		return VersionRollingExtension
	case "minimum-difficulty":
		return MinimumDifficultyExtension
	case "subscribe-extranonce":
		return SubscribeExtranonceExtension
	case "info":
		return InfoExtension
	default:
		return UnknownExtension
	}
}
