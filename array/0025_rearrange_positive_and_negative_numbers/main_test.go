package main

import (
	"reflect"
	"testing"
)

func TestRearrange(t *testing.T) {
	inputArr := []int{1, -1, 2, -3, -3, 0}
	expectedArr := []int{-1, -3, -3, 0, 1, 2}

	Rearrange(inputArr)

	if !reflect.DeepEqual(inputArr, expectedArr) {
		t.Errorf("Expected result: %v, Found: %v", expectedArr, inputArr)
	}
}