package rewe

import (
	"io"

	log "github.com/sirupsen/logrus"
)

type AugmentedBill struct {
	OrderDate          string              `json:"order_date"`
	AugmentedPositions []AugmentedPosition `json:"augmented_positions"`
}

type AugmentedPosition struct {
	Text       string   `json:"text"`
	Categories []string `json:"categories"`
	Amount     int      `json:"amount"`
	Price      Cents    `json:"price"`
	Sum        Cents    `json:"sum"`
	Tax        string   `json:"tax"`
}

func AugmentBill(rs io.ReadSeeker, br BillReader, fetcher CategoryFetcher) (AugmentedBill, error) {
	bill, err := br.Read(rs)
	if err != nil {
		return AugmentedBill{}, err
	}

	var infos []AugmentedPosition
	for _, position := range bill.Positions {
		info, err := fetcher.Fetch(position.Text)
		if err != nil {
			log.Errorf("got err %q - ignoring position %q", err, position.Text)
			continue
		}

		infos = append(infos, AugmentedPosition{
			Text:       info.Product,
			Categories: info.Categories,
			Amount:     position.Amount,
			Price:      position.Price,
			Sum:        position.Sum,
			Tax:        position.Tax,
		})
	}

	return AugmentedBill{
		OrderDate:          bill.OrderDate,
		AugmentedPositions: infos,
	}, nil
}
