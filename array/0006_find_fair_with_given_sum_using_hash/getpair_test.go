package main

import "testing"


func TestGetPairWithSumUsingHash(t *testing.T) {
	inputArr := []int{1, 5, 6, 7, 0}
	sum := 13

	i, j := GetPairWithSumUsingHash(inputArr, sum)

	if i == -1 || j == -1 {
		t.Errorf("Failed to get expected result")
	}
}
