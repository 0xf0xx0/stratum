package stratum_test

import (
	"testing"

	"git.0xf0xx0.eth.limo/0xf0xx0/stratum"
)

/// TODO: tests for each message

func TestMiningSubscribe(t *testing.T) {
	r := makeRequest(`{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-opt-24.5-x64L"]}`)
	if r.GetMethod() != stratum.MethodMiningSubscribe {
		t.Errorf("method mismatch: %s", r.Method)
	}
	s := stratum.MiningSubscribeParams{}
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

	if n.GetMethod() != stratum.MethodClientShowMessage {
		t.Errorf("method mismatch: %s", n.Method)
	}
	s := stratum.ClientShowMessageParams{}
	s.Read(n)

	if s.Message != "Pool restarting; please reconnect." {
		t.Fatalf("message mismatch: %s", s.Message)
	}
}

func TestClientReconnect(t *testing.T) {
	n := makeNotification(`{"id": 0, "method": "client.reconnect", "params":["stratum-lb-usa48.btcguild.com",3333,0]}`)
	if n.GetMethod() != stratum.MethodClientReconnect {
		t.Errorf("method mismatch: %s", n.Method)
	}
	s := stratum.ClientReconnectParams{}
	if err := s.Read(n); err != nil {
		t.Fatalf("read error: %s", err)
	}


	if s.Hostname != "stratum-lb-usa48.btcguild.com" {
		t.Fatal("invalid hostname")
	}
	if s.Port != 3333 {
		t.Log(s.Port)
		t.Fatal("invalid port")
	}
	if s.Waittime != 0 {
		t.Fatal("invalid waittime")
	}
}

func makeRequest(msg string) *stratum.Request {
	r := &stratum.Request{}
	r.UnmarshalJSON([]byte(msg))
	return r
}
func makeNotification(msg string) *stratum.Notification {
	n := &stratum.Notification{}
	n.UnmarshalJSON([]byte(msg))
	return n
}
