package scramble

import (
	"testing"
)

var scrambleTestData = []struct {
	given  string
	expected string
}{
	{"Hello World", "WeHodo lllr"},
	{"john.doe@company.test", "opon.ocdtehsa.t@mjeyn"},
}

func TestScramble(t *testing.T) {
	for _, testData := range scrambleTestData {
		t.Run(testData.given, func(t *testing.T) {
			got := Scramble(testData.given)
			if got != testData.expected {
				t.Errorf("Scramble('%s') = %s, expected %s", testData.given, got, testData.expected)
			}
		})
	}
}
