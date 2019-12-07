package main

import (
	"bufio"
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

func readTransaction(line string) Transaction {
	t := make(Transaction, len(line))
	for i, c := range line {
		t[i] = c != '0'
	}
	return t
}

func readFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			fmt.Println(readTransaction(sc.Text()))
		}
		if err := sc.Err(); err != nil {
			fmt.Println("scan file error:", err)
		}
		f.Close()
	} else {
		fmt.Println("open file error:", err)
	}
}

func main() {
	threshold := flag.Float64("t", 50.0, "threshold to get products")
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("File with transactions should be provided")
		os.Exit(1)
	}

	fmt.Println("threashold=", *threshold)
	fmt.Println("file=", filename)

	readFile(filename)
}
