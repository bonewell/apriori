package main

import "testing"

func (f frequencies) equal(other frequencies) bool {
	if len(f) != len(other) {
		return false
	}
	for i, v := range f {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (f frequencies) notEqual(other frequencies) bool {
	return !f.equal(other)
}

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

func TestParseTransaction(t *testing.T) {
	tran := Parse("0100110")

	if tran.notEqual(Transaction{false, true, false, false, true, true, false}) {
		t.Error(tran)
	}
}

func TestParseEmptyTransaction(t *testing.T) {
	tran := Parse("")

	if len(tran) > 0 {
		t.Error(tran)
	}
}

func (g Goods) equal(other Goods) bool {
	if len(g) != len(other) {
		return false
	}
	for i, v := range g {
		if v != other[i] {
			return false
		}
	}
	return true
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

func TestCount(t *testing.T) {
	goods := GoodsSet{Goods{0}, Goods{1}, Goods{2}, Goods{3}}
	trans := []Transaction{
		{true, true, false, false},
		{true, false, true, false},
		{false, false, false, false},
	}

	freqs := goods.count(trans)

	if freqs.notEqual(frequencies{2, 1, 1, 0}) {
		t.Error(freqs)
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
