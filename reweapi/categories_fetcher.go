package reweapi

import (
	"io"
	"rewe"
	"rewe/util/check"
)

type ReweClient interface {
	GetSearchPage(productName string) (io.Reader, error)
}

type SearchPageParser interface {
	Parse(r io.Reader) (*SearchPage, error)
}

type CategoriesFetcher struct {
	ReweClient       ReweClient
	SearchPageParser SearchPageParser
}

func (c *CategoriesFetcher) Fetch(productName string) (rewe.Categories, error) {
	reader, err := c.ReweClient.GetSearchPage(productName)
	if err != nil {
		return nil, err
	}

	result, err := c.SearchPageParser.Parse(reader)
	if err != nil {
		return nil, err
	}
	check.Equal(len(result.Products), 1, "expected exactly one product")

	return result.Products[0].Categories, nil
}
