package ds

import "testing"

func TestGetIntegerDigitSum(t *testing.T)  {

	value := 1234

	res := GetIntegerDigitSum(value)
	expectedResult := 10

	if res != expectedResult {
		t.Errorf("Result: %d, expected result: %d", res, expectedResult)
	}
}
