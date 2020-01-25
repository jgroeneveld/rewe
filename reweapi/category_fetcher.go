package reweapi

import (
	"errors"

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

	if len(result.Products) < 1 {
		logger.Errorf("no products found for %q", productName)
		return rewe.CategoryInfo{}, ErrNoProductsFound
	}

	if len(result.Products) > 1 {
		logger.Warnf("Got more than 1 product for %q - choosing first %q", productName, result.Products[0].Product)
	}

	return result.Products[0], nil
}

type ReweClient interface {
	GetSearchPage(productName string) (io.Reader, error)
}

type SearchPageParser interface {
	Parse(r io.Reader) (SearchPage, error)
}

var ErrNoProductsFound = errors.New("no products found")
