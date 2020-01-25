package rewe

import (
	"io"

	log "github.com/sirupsen/logrus"
)

type FullProductInfo struct {
	Product    string   `json:"product"`
	Categories []string `json:"categories"`
	Amount     int      `json:"amount"`
	Price      Cents    `json:"price"`
	Sum        Cents    `json:"sum"`
	Tax        string   `json:"tax"`
}

func FetchCategoriesForBill(rs io.ReadSeeker, br BillReader, fetcher CategoryFetcher) ([]FullProductInfo, error) {
	bill, err := br.Read(rs)
	if err != nil {
		return nil, err
	}

	var infos []FullProductInfo
	for _, position := range bill.Positions {
		info, err := fetcher.Fetch(position.Text)
		if err != nil {
			log.Errorf("got err %q - ignoring position %q", err, position.Text)
			continue
		}

		infos = append(infos, FullProductInfo{
			Product:    info.Product,
			Categories: info.Categories,
			Amount:     position.Amount,
			Price:      position.Price,
			Sum:        position.Sum,
			Tax:        position.Tax,
		})
	}

	return infos, nil
}
