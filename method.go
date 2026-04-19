package stratum

import (
	"errors"
)

type Method uint8

const (
	MethodUnknown Method = iota
	MethodClientGetVersion
	MethodClientReconnect
	MethodClientShowMessage
	MethodMiningAuthorize
	MethodMiningConfigure
	MethodMiningExtranonceSubscribe
	MethodMiningNotify
	MethodMiningPing
	MethodMiningSetDifficulty
	MethodMiningSetExtraNonce
	MethodMiningSetVersionMask
	MethodMiningSubmit
	MethodMiningSubscribe
	MethodMiningSuggestDifficulty
)

// String returns the string representation of the [Method].
func (m Method) String() string {
	str, _ := EncodeMethod(m)
	return str
}

// EncodeMethod converts a [Method] to its string name.
func EncodeMethod(m Method) (string, error) {
	switch m {
	case MethodMiningAuthorize:
		return "mining.authorize", nil
	case MethodMiningConfigure:
		return "mining.configure", nil
	case MethodMiningSubscribe:
		return "mining.subscribe", nil
	case MethodMiningExtranonceSubscribe:
		return "mining.extranonce.subscribe", nil
	case MethodMiningNotify:
		return "mining.notify", nil
	case MethodMiningSubmit:
		return "mining.submit", nil
	case MethodMiningSetDifficulty:
		return "mining.set_difficulty", nil
	case MethodMiningSetVersionMask:
		return "mining.set_version_mask", nil
	case MethodMiningSetExtraNonce:
		return "mining.set_extranonce", nil
	case MethodMiningSuggestDifficulty:
		return "mining.suggest_difficulty", nil
	case MethodClientGetVersion:
		return "client.get_version", nil
	case MethodClientReconnect:
		return "client.reconnect", nil
	case MethodMiningPing:
		return "mining.ping", nil
	case MethodClientShowMessage:
		return "client.show_message", nil
	default:
		return "", errors.New("unknown stratum method")
	}
}

// DecodeMethod converts a string name to a [Method].
func DecodeMethod(m string) Method {
	switch m {
	case "mining.authorize":
		return MethodMiningAuthorize
	case "mining.configure":
		return MethodMiningConfigure
	case "mining.subscribe":
		return MethodMiningSubscribe
	case "mining.extranonce.subscribe":
		return MethodMiningExtranonceSubscribe
	case "mining.notify":
		return MethodMiningNotify
	case "mining.submit":
		return MethodMiningSubmit
	case "mining.set_difficulty":
		return MethodMiningSetDifficulty
	case "mining.set_version_mask":
		return MethodMiningSetVersionMask
	case "mining.set_extranonce":
		return MethodMiningSetExtraNonce
	case "mining.suggest_difficulty":
		return MethodMiningSuggestDifficulty
	case "client.get_version":
		return MethodClientGetVersion
	case "client.reconnect":
		return MethodClientReconnect
	case "mining.ping":
		return MethodMiningPing
	case "client.show_message":
		return MethodClientShowMessage
	default:
		return MethodUnknown
	}
}
