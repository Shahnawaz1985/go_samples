package unittest

import (
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Errorf("didn't match expected values")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "watermelon"}
	if JoinWithCommas(list) != "apple, orange, and watermelon" {
		t.Errorf("didn't match expected values")
	}
}
