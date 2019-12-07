package main

import (
	"flag"
	"fmt"
	"os"
)

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
	threshold := flag.Float64("t", 50.0, "threshold to get products")
	flag.Parse()
	inputFile := flag.Arg(0)
	if inputFile == "" {
		fmt.Println("File with transactions should be provided")
		os.Exit(1)
	}

	fmt.Println("threashold=", *threshold)
	fmt.Println("file=", inputFile)
}
