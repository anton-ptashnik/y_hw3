package main

import (
	"reflect"
	"testing"
)

var testData = []struct {
	input  [][]int
	output []int
	err    bool
}{
	{[][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	{[][]int{
		{1, 2},
		{4, 3},
	}, []int{1, 2, 3, 4}, false},
	{[][]int{{1}}, []int{1}, false},
	{[][]int{}, []int{}, false},
	{[][]int{
		{1, 2},
		{4},
	}, []int{1, 2, 3, 4}, true},
}

func TestFlatTable(t *testing.T) {
	for _, d := range testData {
		actual, err := flatTable(d.input)
		if err == nil != d.err {
			continue
		}
		if !reflect.DeepEqual(actual, d.output) {
			t.Errorf("err: %v, input: %v\nactual: %v, expected: %v\n", err, d.input, actual, d.output)
		}
	}
}
