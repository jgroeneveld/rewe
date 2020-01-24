package rewe

import (
	"io"
)

func FetchCategoriesForBill(rs io.ReadSeeker, br BillReader, fetcher CategoryFetcher) ([]CategoryInfo, error) {
	bill, err := br.Read(rs)
	if err != nil {
		return nil, err
	}

	var infos []CategoryInfo
	for _, position := range bill.Positions {
		info, err := fetcher.Fetch(position.Text)
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}
