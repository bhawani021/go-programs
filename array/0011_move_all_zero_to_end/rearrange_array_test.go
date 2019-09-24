package main

import (
	"reflect"
	"testing"
)

func TestArrayZeroToEnd(t *testing.T) {
	inputArr := []int{1, 0, 2, 0}
	expectedArray := []int{1, 2, 0, 0}
	ArrayZeroToEnd(inputArr)

	if !reflect.DeepEqual(inputArr, expectedArray) {
		t.Errorf("Expected result: %v, found: %v", expectedArray, inputArr)
	}
}
