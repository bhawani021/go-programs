package main

import "testing"


func TestSearch(t *testing.T) {
	inputArr := []int{10, 12, 14, 0, 1, 2, 3}

	result := Search(inputArr)
	expectedResult := 3 // Index of minimum value

	if result != expectedResult {
		t.Errorf("Expected result: %d, found: %d", expectedResult, result)
	}
}