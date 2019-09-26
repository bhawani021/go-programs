package main

import (
	"reflect"
	"testing"
)

func TestReArrangeArray(t *testing.T) {
	inputArray := []int{0, 2, 2, 2, 0}
	expectedOutput := []int{4, 2, 0, 0, 0}

	ReArrangeArray(inputArray)

	if !reflect.DeepEqual(inputArray, expectedOutput) {
		t.Errorf("Expected output: %v. Found: %v", expectedOutput, inputArray)
	}
}
