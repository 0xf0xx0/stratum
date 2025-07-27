package Stratum

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
		return "", errors.New("unknown Stratum extension")
	}
}

func DecodeExtension(m string) (Extension, error) {
	switch m {
	case "version-rolling":
		return VersionRolling, nil
	case "minimum-difficulty":
		return MinimumDifficulty, nil
	case "subscribe-extranonce":
		return SubscribeExtranonce, nil
	case "info":
		return Info, nil
	default:
		return Unknown, errors.New("unknown Stratum extension")
	}
}
