package ds

import "testing"

func TestGetPairWithGivenSumInArr(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}

	sum := 70

	i, j := GetPairWithGivenSumInArr(arr, sum)

	// Expected result
	iExpected, jExpected := 1, 4

	if i != iExpected || j != jExpected {
		t.Errorf("Received expected: %d, %d. Expected result: %d %d",
			i, j, iExpected, jExpected)
	}
}
