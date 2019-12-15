package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Transaction []bool // ordered by index of the product
type Goods []int
type GoodsSet []Goods

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

func (g Goods) equal(o Goods) bool {
	if len(g) != len(o) {
		return false
	}
	for i, v := range g {
		if v != o[i] {
			return false
		}
	}
	return true
}

func (g Goods) union(o Goods) Goods {
	goods := Goods{}
	i, j := 0, 0
	for i < len(g) && j < len(o) {
		if g[i] < o[j] {
			goods = append(goods, g[i])
			i++
		} else if g[i] > o[j] {
			goods = append(goods, o[j])
			j++
		} else {
			goods = append(goods, g[i])
			i++
			j++
		}
	}
	var tail Goods
	if i < j {
		tail = g[i:]
	} else {
		tail = o[j:]
	}
	return append(goods, tail...)
}

func (gs GoodsSet) init() GoodsSet {
	for i := 0; i < len(gs); i++ {
		gs[i] = Goods{i}
	}
	return gs
}

func (gs GoodsSet) generate(k int) GoodsSet {
	gen := GoodsSet{}
	for i := 0; i < len(gs); i++ {
		for j := i; j < len(gs); j++ {
			gsi := gs[i][:k-2]
			gsj := gs[j][:k-2]
			if gsi.equal(gsj) {
				goods := gs[i].union(gs[j])
				if len(goods) == k {
					gen = append(gen, goods)
				}
			}
		}
	}
	return gen
}

func (gs GoodsSet) prune(transactions []Transaction, threshold int) GoodsSet {
	res := GoodsSet{}
	fs := make([]int, len(gs))
	for i, goods := range gs {
		for _, t := range transactions {
			if t.contains(goods) {
				fs[i]++
			}
		}
		if fs[i] >= threshold {
			res = append(res, gs[i])
		}
	}
	return res
}

func Apriori(transactions []Transaction, threshold int) GoodsSet {
	if len(transactions) > 0 {
		count := len(transactions[0])
		res := GoodsSet{}
		gs := make(GoodsSet, count).init().prune(transactions, threshold)
		res = append(res, gs...)
		for k := 2; len(gs) > 0; k++ {
			gs = gs.generate(k).prune(transactions, threshold)
			res = append(res, gs...)
		}
		return res
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
