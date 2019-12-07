package main

import "fmt"

type Transaction []bool
type WrongTransaction string
type Frequencies []int

func (t WrongTransaction) Error() string {
	return string(t)
}

func (f Frequencies) update(t Transaction) error {
	if len(f) != len(t) {
		return WrongTransaction("Count of products in transaction is not correct")
	}
	for i, presented := range t {
		if presented {
			f[i] += 1
		}
	}
	return nil
}

func main() {
	fmt.Printf("Apriori goes!")
}
