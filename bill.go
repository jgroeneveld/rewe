package rewe

import "io"

type Bill struct {
	OrderDate string     `json:"order_date"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Text   string `json:"text"`
	Amount int    `json:"amount"`
	Price  Cents  `json:"price"`
	Sum    Cents  `json:"sum"`
	Tax    string `json:"tax"`
}

type Cents int

//go:generate mockgen -source=bill.go -package=rewe -destination mock_bill_test.go
type BillReader interface {
	Read(io.ReadSeeker) (Bill, error)
}
