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
	size := len(t)
	for _, index := range goods {
		if size <= index || !t[index] {
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
	sizeg, sizeo := len(g), len(o)
	goods := make(Goods, 0, sizeg + sizeo)
	i, j := 0, 0
	for i < sizeg && j < sizeo {
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
	size := len(gs)
	gen := make(GoodsSet, 0, 2 * size)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
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

func (g Goods) count(transactions []Transaction, frequency *int, c chan bool) {
	for _, t := range transactions {
		if t.contains(g) {
			*frequency++
		}
	}
	c <- true
}

func (gs GoodsSet) prune(transactions []Transaction, threshold float64) GoodsSet {
	size := len(gs)
	count := float64(len(transactions))
	res := make(GoodsSet, 0, size)
	c := make(chan bool)
	fs := make([]int, size)
	for i, goods := range gs {
		go goods.count(transactions, &fs[i], c)
	}
	for _ = range gs {
		<- c
	}
	for i, goods := range gs {
		if float64(fs[i])/count >= threshold {
			res = append(res, goods)
		}
	}
	return res
}

func Apriori(transactions []Transaction, threshold float64) GoodsSet {
	if len(transactions) > 0 {
		count := len(transactions[0])
		res := make(GoodsSet, 0, 2 * count)
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
	threshold := flag.Float64("t", 0.5, "threshold to get products")
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
