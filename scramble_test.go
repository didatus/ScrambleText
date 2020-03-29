package scramble

import (
	"testing"
)

func TestScramble(t *testing.T) {
	got := Scramble("Hello World")
	if got != "WeHodo lllr" {
		t.Errorf("Hello() = %s, want %s", got, "x")
	}
}

func TestScramble2(t *testing.T) {
	got := Scramble("john.doe@company.test")
	if got != "opon.ocdtehsa.t@mjeyn" {
		t.Errorf("Hello() = %s, want %s", got, "x")
	}
}
