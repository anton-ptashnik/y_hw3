package main

import (
	"reflect"
	"testing"
)

var testData = []struct {
	input  [][]int
	output []int
}{
	{[][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[][]int{{1}}, []int{1}},
	{[][]int{}, []int{}},
}

func TestFlatTable(t *testing.T) {
	for _, d := range testData {
		actual := flatTable(d.input)
		if !reflect.DeepEqual(actual, d.output) {
			t.Errorf("input: %v\nactual: %v, expected: %v\n", d.input, actual, d.output)
		}
	}
}
