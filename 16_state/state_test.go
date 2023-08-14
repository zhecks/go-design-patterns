package state

import (
	"testing"
)

func TestState(t *testing.T) {
	manager := NewVoteManager()
	for i := 0; i < 8; i++ {
		manager.Vote("Alice", "item")
	}
}
