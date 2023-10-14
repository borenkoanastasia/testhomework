package tests

import (
	"testhomework/hw2/pkg"
	"testing"
)

func TestHw2English(t *testing.T) {
	expected := "Hello"
	actual := pkg.English()

	if expected != actual {
		t.Errorf("Result  got: %s, want: %s.", actual, expected)
	}
}
