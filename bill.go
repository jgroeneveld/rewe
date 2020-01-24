package rewe

import "io"

type Bill struct {
	Positions []Position
}

type Position struct {
	Text   string
	Amount int
	Price  Cents
	Sum    Cents
	Tax    string
}

type Cents int

type BillReader interface {
	Read(io.ReadSeeker) (Bill, error)
}
