package tests

import (
	"testhomework/hw3/pkg"
	"testing"
)

func TestHw3English(t *testing.T) {
	expected := "Hello Bean"
	actual := pkg.English()

	if expected != actual {
		t.Errorf("Result  got: %s, want: %s.", actual, expected)
	}
}
