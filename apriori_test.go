package main

import "testing"

func (f Frequencies) equal(other Frequencies) bool {
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

func (f Frequencies) notEqual(other Frequencies) bool {
	return !f.equal(other)
}

func TestCalculateFrequency(t *testing.T) {
	fs := make(Frequencies, 4)
	fs.update(Transaction{true, true, false, true})
	fs.update(Transaction{false, true, false, true})
	fs.update(Transaction{false, true, false, true})
	fs.update(Transaction{true, false, false, false})
	fs.update(Transaction{false, false, false, true})

	if fs.notEqual(Frequencies{2, 3, 0, 4}) {
		t.Error(fs)
	}
}

func TestNoTransactions(t *testing.T) {
	fs := make(Frequencies, 5)

	if fs.notEqual(Frequencies{0, 0, 0, 0, 0}) {
		t.Error(fs)
	}
}

func TestWrongTransaction(t *testing.T) {
	fs := make(Frequencies, 3)

	if err := fs.update(Transaction{true, false}); err == nil {
		t.Fail()
	}
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

func TestReadTransaction(t *testing.T) {
	tran := readTransaction("0100110")

	if tran.notEqual(Transaction{false, true, false, false, true, true, false}) {
		t.Error(tran)
	}
}
