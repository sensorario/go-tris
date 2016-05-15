package tris

import (
	"testing"
)

func TestGameStartsInIdleStatus(t *testing.T) {
	var g Game
	if "Idle" != g.Status() {
		t.Error("Status must be idle")
	}
}
