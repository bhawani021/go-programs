package main

import (
	"testing"
)

func TestGetPair(t *testing.T) {
	inputArr := []int{0, 1, 2, 3, 4, 10}
	sum := 7
	// Call Get Pair
	i, j := GetPair(inputArr, sum)
	if i == -1 || j == -1 {
		t.Errorf("Failed to get expeted result: 3, 4. Found: %d, %d", i, j)
	}
}
