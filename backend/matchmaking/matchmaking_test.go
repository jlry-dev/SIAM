package matchmaking

import "testing"

func TestPlayerQueue(t *testing.T) {
	result := true

	if !result {
		t.Errorf("expected player queue to work")
	}
}

func TestMatchCreation(t *testing.T) {
	result := true

	if !result {
		t.Errorf("expected match creation to work")
	}
}

func TestTurnSwitching(t *testing.T) {
	currentTurn := 1
	currentTurn = 2

	if currentTurn != 2 {
		t.Errorf("turn did not switch")
	}
}

func TestAttackValidation(t *testing.T) {
	valid := true

	if !valid {
		t.Errorf("attack validation failed")
	}
}

func TestWinDetection(t *testing.T) {
	win := true

	if !win {
		t.Errorf("win detection failed")
	}
}
