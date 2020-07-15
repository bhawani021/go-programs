package ds

import (
	"testing"
)

func TestIsStringPerfectReversible(t *testing.T) {

	str := "abcxcba"

	res := IsStringPerfectReversible(str)

	if res == false {
		t.Errorf("Got result: %t. Expected result: %t", res, true)
	}
}
