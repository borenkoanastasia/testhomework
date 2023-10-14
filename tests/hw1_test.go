package tests

import (
	"testhomework/hw1/pkg"
	"testing"
)

func TestHw1English(t *testing.T) {
	expected := "Hello"
	actual := pkg.English()

	if expected != actual {
		t.Errorf("Result  got: %s, want: %s.", actual, expected)
	}
}
