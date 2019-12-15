package main

import "testing"

func (t Transaction) equal(other Transaction) bool {
	if len(t) != len(other) {
		return false
	}
	for i, v := range t {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (t Transaction) notEqual(other Transaction) bool {
	return !t.equal(other)
}

func (g Goods) notEqual(other Goods) bool {
	return !g.equal(other)
}

func (gs GoodsSet) equal(other GoodsSet) bool {
	if len(gs) != len(other) {
		return false
	}
	for i, v := range gs {
		if v.notEqual(other[i]) {
			return false
		}
	}
	return true
}

func (gs GoodsSet) notEqual(other GoodsSet) bool {
	return !gs.equal(other)
}

func TestParseTransaction(t *testing.T) {
	tran := Parse("0100110")

	if tran.notEqual(Transaction{false, true, false, false, true, true, false}) {
		t.Error(tran)
	}
}

func BenchmarkParseTransaction(b *testing.B) {
	var _ = Parse("0000000000010000000000000000010000000000001000000000000010000000000000000000000000000000000000000000")
}

func TestParseEmptyTransaction(t *testing.T) {
	tran := Parse("")

	if len(tran) > 0 {
		t.Error(tran)
	}
}

func TestContains(t *testing.T) {
	tran := Transaction{true, false, true}
	if !tran.contains(Goods{0, 2}) {
		t.Error(false)
	}
}

func TestContainsEmptyGoods(t *testing.T) {
	tran := Transaction{true, false, true}
	if !tran.contains(Goods{}) {
		t.Error(false)
	}
}

func TestContainsEmptyTransaction(t *testing.T) {
	tran := Transaction{}
	if tran.contains(Goods{1}) {
		t.Error(true)
	}
}

func BenchmarkContains(b *testing.B) {
	tran := Transaction{
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
		true, false, true, false, false, true, true, false, true, true,
	}
	goods := Goods{1, 2, 5, 7, 11, 24, 35, 41, 55, 62, 73, 87, 99}
	var _ = tran.contains(goods)
}

func TestUnion(t *testing.T) {
	goods := Goods{0}.union(Goods{1})

	if goods.notEqual(Goods{0, 1}) {
		t.Error(goods)
	}
}

func TestUnionSimilar(t *testing.T) {
	goods := Goods{0, 1, 2}.union(Goods{1, 2, 4})

	if goods.notEqual(Goods{0, 1, 2, 4}) {
		t.Error(goods)
	}
}

func TestUnionTail(t *testing.T) {
	goods := Goods{0, 1, 2}.union(Goods{0, 1, 2, 3, 4, 5})

	if goods.notEqual(Goods{0, 1, 2, 3, 4, 5}) {
		t.Error(goods)
	}
}

func TestUnionEmpty(t *testing.T) {
	goods := Goods{}.union(Goods{})

	if len(goods) != 0 {
		t.Error(goods)
	}
}

func BenchmarkGoodsUnion(b *testing.B) {
	goods := Goods{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 12, 13, 14, 17, 18, 19,
		20, 22, 23, 24, 27, 28, 29,
		30, 32, 33, 34, 37, 38, 39,
		40, 42, 43, 44, 47, 48, 49,
		50, 52, 53, 54, 57, 58, 59,
		60, 62, 63, 64, 67, 68, 69,
		70, 72, 73, 74, 77, 78, 79,
		80, 82, 83, 84, 87, 88, 89,
		90, 92, 93, 94, 97, 98, 99,
	}

	goods2 := Goods{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		11, 12, 15, 16, 17, 18, 19,
		21, 22, 25, 26, 27, 28, 29,
		31, 32, 35, 36, 37, 38, 39,
		41, 42, 45, 46, 47, 48, 49,
		51, 52, 55, 56, 57, 58, 59,
		61, 62, 65, 66, 67, 68, 69,
		71, 72, 75, 76, 77, 78, 79,
		81, 82, 85, 86, 87, 88, 89,
		91, 92, 95, 96, 97, 98, 99,
	}

	var _ = goods.union(goods2)
}

func TestGenerateDouble(t *testing.T) {
	gs := GoodsSet{{1}, {2}, {3}}

	gs2 := gs.generate(2)

	if gs2.notEqual(GoodsSet{{1, 2}, {1, 3}, {2, 3}}) {
		t.Error(gs2)
	}
}

func TestGenerateTriple(t *testing.T) {
	gs := GoodsSet{{1, 2}, {1, 3}, {2, 3}}

	gs2 := gs.generate(3)

	if gs2.notEqual(GoodsSet{{1, 2, 3}}) {
		t.Error(gs2)
	}
}

func TestGenerateFours(t *testing.T) {
	gs := GoodsSet{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {1, 3, 5}, {2, 3, 4}}

	gs2 := gs.generate(4)

	if gs2.notEqual(GoodsSet{{1, 2, 3, 4}, {1, 3, 4, 5}}) {
		t.Error(gs2)
	}
}

func BenchmarkGenerate(b *testing.B) {
	gs := GoodsSet{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}

	var _ = gs.generate(2)
}

func TestInitialize(t *testing.T) {
	g := make(GoodsSet, 5).init()

	if g.notEqual(GoodsSet{{0}, {1}, {2}, {3}, {4}}) {
		t.Error(g)
	}
}

func TestInitializeEmpty(t *testing.T) {
	g := make(GoodsSet, 0).init()

	if len(g) != 0 {
		t.Error(g)
	}
}

func TestPrune(t *testing.T) {
	gs := GoodsSet{Goods{0}, Goods{1}, Goods{2}, Goods{3}}
	trans := []Transaction{
		{true, true, false, false},
		{true, false, false, true},
		{false, false, false, false},
	}

	pruned := gs.prune(trans, 0.3)

	if pruned.notEqual(GoodsSet{{0}, {1}, {3}}) {
		t.Error(pruned)
	}
}

func BenchmarkPrune(b *testing.B) {
	gs := GoodsSet{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	trans := []Transaction{
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
	}
	var _ = gs.prune(trans, 7)
}

func TestApriori(t *testing.T) {
	trans := []Transaction{
		{true, false, false, false, true},
		{true, false, false, true, false},
		{true, false, true, false, false},
		{false, false, true, true, false},
		{false, false, true, false, true},
	}

	gs := Apriori(trans, 0.4)
	if gs.notEqual(GoodsSet{{0}, {2}, {3}, {4}}) {
		t.Error(gs)
	}
}

func BenchmarkApriori(b *testing.B) {
	trans := []Transaction{
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
	}

	var _ = Apriori(trans, 0.4)
}

func TestCount(t *testing.T) {
	trans := []Transaction{
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
	}

	freq := 0
	c := make(chan bool)
	go Goods{0, 2}.count(trans, &freq, c)
	<-c

	if freq != 10 {
		t.Error(freq)
	}
}

func BenchmarkCount(b *testing.B) {
	trans := []Transaction{
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
		{true, false, true, false, false, true, true, false, true, true},
	}

	freq := 0
	c := make(chan bool)
	go Goods{0, 2}.count(trans, &freq, c)
	<-c
}
