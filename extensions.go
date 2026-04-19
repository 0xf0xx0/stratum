package stratum

import (
	"errors"
)

// https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki

type Extension uint8

const (
	UnknownExtension Extension = iota
	VersionRollingExtension
	MinimumDifficultyExtension
	SubscribeExtranonceExtension
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
