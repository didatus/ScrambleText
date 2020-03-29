package scramble

import (
	"fmt"
	"testing"
)

func TestScramble(t *testing.T) {
	testData := map[string]string {
		"Hello World": "WeHodo lllr",
		"john.doe@company.test": "opon.ocdtehsa.t@mjeyn",
	}
	for given, expected := range testData {
		fmt.Printf("%s -> %s\n", given, expected)

		got := Scramble(given)
		if got != expected {
			t.Errorf("Scramble('%s') = %s, expected %s", given, got, expected)
		}
	}
}
