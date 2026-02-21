package stratum_test

import (
	"testing"

	"git.0xf0xx0.eth.limo/0xf0xx0/stratum"
)

/// TODO: tests for each message

func TestMiningSubscribe(t *testing.T) {
	r := makeRequest(`{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-opt-24.5-x64L"]}`)
	if r.GetMethod() != stratum.MiningSubscribe {
		t.Errorf("method mismatch: %s", r.Method)
	}
	s := stratum.SubscribeParams{}
	s.Read(r)
	if s.UserAgent != "cpuminer-opt-24.5-x64L" {
		t.Errorf("useragent mismatch: %s", s.UserAgent)
	}
	if r.MessageID != stratum.MessageID(1) {
		t.Errorf("message id mismatch: %d", r.MessageID)
	}
}

func TestClientShowMessage(t *testing.T) {
	n := makeNotification(`{"id":null,"method":"client.show_message","params":["Pool restarting; please reconnect."]}`)

	if n.GetMethod() != stratum.ClientShowMessage {
		t.Errorf("method mismatch: %s", n.Method)
	}
	s := stratum.ShowMessageParams{}
	s.Read(n)

	if s.Message != "Pool restarting; please reconnect." {
		t.Fatalf("message mismatch: %s", s.Message)
	}
}

func TestClientReconnect(t *testing.T) {
	n := makeNotification(`{"id": 0, "method": "client.reconnect", "params":["stratum-lb-usa48.btcguild.com",3333,0]}`)
	if n.GetMethod() != stratum.ClientReconnect {
		t.Errorf("method mismatch: %s", n.Method)
	}
}

func makeRequest(msg string) *stratum.Request {
	r := &stratum.Request{}
	r.Unmarshal([]byte(msg))
	return r
}
func makeNotification(msg string) *stratum.Notification {
	n := &stratum.Notification{}
	n.Unmarshal([]byte(msg))
	return n
}
