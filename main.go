package main

import "hw3/spiraliterator"

func main() {

}

func flatTable(t [][]int) []int {
	r := make([]int, 0, len(t)*len(t))
	var i = spiraliterator.New(t)
	for i.Next() {
		v := i.Read()
		r = append(r, v)
	}
	return r
}
