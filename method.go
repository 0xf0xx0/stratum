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

func (m Method) String() string {
	str, _ := EncodeMethod(m)
	return str
}

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
		return "mining.set_extra_nonce", nil
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

func DecodeMethod(m string) (Method, error) {
	switch m {
	case "mining.authorize":
		return MethodMiningAuthorize, nil
	case "mining.configure":
		return MethodMiningConfigure, nil
	case "mining.subscribe":
		return MethodMiningSubscribe, nil
	case "mining.extranonce.subscribe":
		return MethodMiningExtranonceSubscribe, nil
	case "mining.notify":
		return MethodMiningNotify, nil
	case "mining.submit":
		return MethodMiningSubmit, nil
	case "mining.set_difficulty":
		return MethodMiningSetDifficulty, nil
	case "mining.set_version_mask":
		return MethodMiningSetVersionMask, nil
	case "mining.set_extra_nonce":
		return MethodMiningSetExtraNonce, nil
	case "mining.suggest_difficulty":
		return MethodMiningSuggestDifficulty, nil
	case "client.get_version":
		return MethodClientGetVersion, nil
	case "client.reconnect":
		return MethodClientReconnect, nil
	case "mining.ping":
		return MethodMiningPing, nil
	case "client.show_message":
		return MethodClientShowMessage, nil
	default:
		return MethodUnknown, errors.New("unknown stratum method")
	}
}
