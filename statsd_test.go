package statsd

import (
	"testing"
)

func TestTrySendLocalhost(t *testing.T) {
	client := New("localhost", 8125)

	client.TrySend(&Metric{
		Name:  "test",
		Value: 1,
		Type:  Counter,
	})
}

func TestTrySendToAInvalidAddress(t *testing.T) {
	client := New("invalid.address", 8125)

	client.TrySend(&Metric{
		Name:  "test",
		Value: 1,
		Type:  Counter,
	})
}
