package rewebill

import "io"

func Read(rs io.ReadSeeker) (Bill, error) {
	pdf, err := ReadPdf(rs)
	if err != nil {
		return Bill{}, err
	}

	bill, err := Extract(pdf)
	if err != nil {
		return Bill{}, err
	}

	return bill, nil
}
