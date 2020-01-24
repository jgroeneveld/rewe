package rewebill

import (
	"io"
	"rewe"
)

var Reader = readerFunc(Read)

type readerFunc func(rs io.ReadSeeker) (rewe.Bill, error)

func (fn readerFunc) Read(rs io.ReadSeeker) (rewe.Bill, error) {
	return fn(rs)
}

func Read(rs io.ReadSeeker) (rewe.Bill, error) {
	pdf, err := ReadPdf(rs)
	if err != nil {
		return rewe.Bill{}, err
	}

	bill, err := Extract(pdf)
	if err != nil {
		return rewe.Bill{}, err
	}

	return bill, nil
}
