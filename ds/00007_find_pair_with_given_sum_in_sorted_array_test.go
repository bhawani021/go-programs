package ds

import "testing"

func TestFindPairWithGivenSumInSortedArray(t *testing.T) {
	arr := []int{1, 2, 3, 4}

	i, j := FindPairWithGivenSumInSortedArray(arr, 7)

	iExpected, jExpected := 2, 3

	if i != iExpected || j != jExpected {
		t.Errorf("Result: %d %d, expected result: %d %d", i, j, iExpected, jExpected)
	}
}
