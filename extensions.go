package stratum

import (
	"errors"
)

// https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki

type Extension uint8

const (
	Unknown = iota
	VersionRolling
	MinimumDifficulty
	SubscribeExtranonce
	Info
)

func EncodeExtension(m Extension) (string, error) {
	switch m {
	case VersionRolling:
		return "version-rolling", nil
	case MinimumDifficulty:
		return "minimum-difficulty", nil
	case SubscribeExtranonce:
		return "subscribe-extranonce", nil
	case Info:
		return "info", nil
	default:
		return "", errors.New("unknown stratum extension")
	}
}

func DecodeExtension(m string) Extension {
	switch m {
	case "version-rolling":
		return VersionRolling
	case "minimum-difficulty":
		return MinimumDifficulty
	case "subscribe-extranonce":
		return SubscribeExtranonce
	case "info":
		return Info
	default:
		return Unknown
	}
}
