package main

import "fmt"

func CalculateFrequency(transactions [][]int) []int {
	frequency := make([]int, len(transactions[0]))
	for _, t := range transactions {
		for i, f := range t {
			if f == 1 {
				frequency[i] += 1
			}
		}
	}
	return frequency
}

func main() {
	fmt.Printf("Apriori goes!")
}
