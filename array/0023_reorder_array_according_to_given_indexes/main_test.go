package main

import (
	"reflect"
	"testing"
)

func TestRearrange(t *testing.T) {

	arr := []int{50, 40, 70, 60, 90}
	index := []int{3,  0,  4,  1,  2}

	expectedArr := []int{40, 60, 90, 50, 70}
	expectedIndex := []int{0, 1, 2, 3, 4}

	Rearrange(arr, index)

	if !reflect.DeepEqual(arr, expectedArr) {
		t.Errorf("Expeced result: %v. Found: %v", expectedArr, arr)
	}

	if !reflect.DeepEqual(expectedIndex, index) {
		t.Errorf("Expected result: %v. found: %v", expectedIndex, index)
	}
}
