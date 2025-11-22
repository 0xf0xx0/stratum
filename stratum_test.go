package stratum_test

import (
	"testing"

	"github.com/0xf0xx0/stratum"
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

func BenchmarkRequest(b *testing.B) {
	for b.Loop() {
		r := makeRequest(`{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-opt-24.5-x64L"]}`)
		s := stratum.SubscribeParams{}
		s.Read(r)
		r.Respond(true)
		_,_=r.Marshal()
	}
}

func makeRequest(msg string) *stratum.Request {
	r := &stratum.Request{}
	r.Unmarshal([]byte(msg))
	return r
}
