package main

import (
	"reflect"
	"testing"
)

func TestSegregate(t *testing.T) {
	inputArr := []int{1, 0, -1, 0}

	expectedOutput := []int{1, 0, 0, -1}

	Segregate(inputArr)

	if !reflect.DeepEqual(inputArr, expectedOutput) {
		t.Errorf("Expected result: %v. found: %v", expectedOutput, inputArr)
	}


}
