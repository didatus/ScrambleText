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
	{"你好，世界", "你，好世界"},
}

func TestScramble(t *testing.T) {
	t.Parallel()
	for _, testData := range scrambleTestData {
		testData := testData
		t.Run(testData.given, func(t *testing.T) {
			got := Scramble(testData.given)
			if got != testData.expected {
				t.Parallel()
				t.Errorf("Scramble('%s') = %s, expected %s", testData.given, got, testData.expected)
			}
		})
	}
}
