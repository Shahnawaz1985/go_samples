package unittest

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	} else {
		fmt.Println(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "watermelon"}
	want := "apple, orange, and watermelon"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	} else {
		fmt.Println(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas (%#v) = \"%s\", want \"%s\"", list, got, want)
}
