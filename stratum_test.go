package stratum_test

import (
	"testing"

	"github.com/0xf0xx0/stratum"
)

/// TODO: tests for each message

func TestMiningSubscribe(t *testing.T) {
	r := makeRequest(`{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-opt-24.5-x64L"]}`)
	m, _ := stratum.EncodeMethod(stratum.MiningSubscribe)
	if r.Method != m {
		t.Errorf("method mismatch: %s", r.Method)
	}
	s := stratum.SubscribeParams{}
	s.Read(r)
	if s.UserAgent != "cpuminer-opt-24.5-x64L" {
		t.Errorf("useragent mismatch: %s", s.UserAgent)
	}
	// TODO: literally why doesnt thi work
	if r.MessageID != stratum.MessageID(1) {
		t.Errorf("message id mismatch: %d", r.MessageID)
	}
}

func makeRequest(msg string) *stratum.Request {
	r := &stratum.Request{}
	r.Unmarshal([]byte(msg))
	return r
}
