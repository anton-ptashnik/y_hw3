package main

import "hw3/spiraliterator"

func flatTable(t [][]int) ([]int, error) {
	r := make([]int, 0, len(t)*len(t))
	var i, err = spiraliterator.Create(t)
	for i.Next() {
		v := i.Read()
		r = append(r, v)
	}
	return r, err
}
