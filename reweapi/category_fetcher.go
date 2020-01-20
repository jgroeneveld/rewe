package reweapi

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"io"
	"rewe"
)

type CategoryFetcher struct {
	ReweClient       ReweClient
	SearchPageParser SearchPageParser
}

func (c CategoryFetcher) Fetch(productName string) (rewe.CategoryInfo, error) {
	logger := log.WithField("Caller", "CategoriesFetcher.Fetch")
	logger.Infof("Fetching categories %q", productName)

	reader, err := c.ReweClient.GetSearchPage(productName)
	if err != nil {
		return rewe.CategoryInfo{}, err
	}

	result, err := c.SearchPageParser.Parse(reader)
	if err != nil {
		return rewe.CategoryInfo{}, err
	}
	if len(result.Products) != 1 {
		return rewe.CategoryInfo{}, &ErrFuzzyResult{Products: result.Products}
	}

	return result.Products[0], nil
}

type ReweClient interface {
	GetSearchPage(productName string) (io.Reader, error)
}

type SearchPageParser interface {
	Parse(r io.Reader) (SearchPage, error)
}

type ErrFuzzyResult struct {
	Products []rewe.CategoryInfo
}

func (err ErrFuzzyResult) Error() string {
	return fmt.Sprintf("Fuzzy Result. Expected one product got %d", len(err.Products))
}
