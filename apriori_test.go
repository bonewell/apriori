package main

import "testing"

func TestCalculateFrequency(t *testing.T) {
	trans := [][]int{
		[]int{1, 1, 0, 1},
		[]int{0, 1, 0, 1},
		[]int{0, 0, 0, 1},
		[]int{1, 1, 0, 0},
		[]int{0, 0, 0, 1},
	}

	frequency := CalculateFrequency(trans)

	if len(frequency) != 4 {
		t.Fatal(len(frequency))
	}
	var expected []int = []int{2, 3, 0, 4}
	for i, f := range frequency {
		if f != expected[i] {
			t.Error(i, f)
		}
	}
}
