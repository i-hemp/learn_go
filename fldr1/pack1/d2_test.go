package pack1

import (
	// "fldr1/pack1"
	"testing"
)

func TestDev(t *testing.T) {
	if Dev() != "ev" {
		t.Fatal("wrong :(")
	}
}
