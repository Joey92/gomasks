package gomasks

import "testing"

const (
	TESTFLAG_ONE Bitmask = 1 << iota
	TESTFLAG_TWO
	TESTFLAG_THREE
)

func TestAddFlag(t *testing.T) {

	var mainFlag Bitmask = TESTFLAG_TWO

	mainFlag.AddFlag(TESTFLAG_THREE)

	if mainFlag&TESTFLAG_THREE == 0 {
		t.FailNow()
	}

}

func TestRemoveFlag(t *testing.T) {

	var mainFlag Bitmask = TESTFLAG_ONE | TESTFLAG_THREE

	mainFlag.RemoveFlag(TESTFLAG_THREE)

	if mainFlag&TESTFLAG_THREE != 0 {
		t.FailNow()
	}

}

func TestHasFlag(t *testing.T) {

	var mainFlag Bitmask = TESTFLAG_ONE | TESTFLAG_THREE

	if !mainFlag.HasFlag(TESTFLAG_THREE) {
		t.FailNow()
	}

}

func TestToggleFlag(t *testing.T) {

	var mainFlag Bitmask = TESTFLAG_ONE

	if mainFlag&TESTFLAG_THREE != 0 {
		t.FailNow()
	}

	mainFlag.ToggleFlag(TESTFLAG_THREE)

	if mainFlag&TESTFLAG_THREE == 0 {
		t.FailNow()
	}

	mainFlag.ToggleFlag(TESTFLAG_THREE)

	if mainFlag&TESTFLAG_THREE != 0 {
		t.FailNow()
	}
}
