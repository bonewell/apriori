package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Transaction []bool  // ordered by index of the product
type Goods []int
type GoodsSet []Goods
type frequencies []int

func Parse(line string) Transaction {
	t := make(Transaction, len(line))
	for i, c := range line {
		t[i] = c != '0'
	}
	return t
}

func load(filename string) []Transaction {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err == nil {
		transactions := make([]Transaction, 0)
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			transactions = append(transactions, Parse(sc.Text()))
		}
		if err := sc.Err(); err != nil {
			fmt.Println("scan file error:", err)
		}
		f.Close()
		return transactions
	} else {
		fmt.Println("open file error:", err)
		return []Transaction{}
	}
}

func (t Transaction) contains(goods Goods) bool {
	for _, index := range goods {
		if len(t) <= int(index) || !t[index] {
			return false
		}
	}
	return true
}

func (gs GoodsSet) count(ts []Transaction) frequencies {
	fs := make(frequencies, len(gs))
	for i, goods := range gs {
		for _, t := range ts {
			if t.contains(goods) {
				fs[i] += 1
			}
		}
	}
	return fs
}

func (f frequencies) filter(gs GoodsSet, threshold int) GoodsSet {
	if len(f) != len(gs) { panic("Wrong length of data") }
	filtered := make(GoodsSet, 0)
	for i, v := range f {
		if v >= threshold {
			filtered = append(filtered, gs[i])
		}
	}
	return filtered
}

func Apriori(transactions []Transaction, threshold int) GoodsSet {
	if len(transactions) > 0 {
		fmt.Println(len(transactions[0]))
		gs := make(GoodsSet, len(transactions[0]))
		for i := range transactions[0] {
			gs[i] = Goods{i}
		}
		fs := gs.count(transactions)
		return fs.filter(gs, threshold)
	}
	return GoodsSet{}
}

func main() {
	threshold := flag.Int("t", 20, "threshold to get products")
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("File with transactions should be provided")
		os.Exit(1)
	}

	fmt.Println("threashold=", *threshold)
	fmt.Println("file=", filename)

	goods := Apriori(load(filename), *threshold)

	fmt.Println(goods)
}
