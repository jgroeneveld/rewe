package rewe

import "io"

type Bill struct {
	OrderDate string
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

//go:generate mockgen -source=bill.go -package=rewe -destination mock_bill_test.go
type BillReader interface {
	Read(io.ReadSeeker) (Bill, error)
}
