package reweapi

import (
	"io"

	"rewe"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestCategoriesFetcher(t *testing.T) {
	t.Run("returns the parse result", func(t *testing.T) {
		parseResult := SearchPage{[]Product{{"Apfelsaft", rewe.Categories{"saft"}}}}
		fetcher := mockedCategoriesFetcher(t, parseResult)

		categories, err := fetcher.Fetch("Apfelsaft")
		assert.NilError(t, err)

		assert.DeepEqual(t, categories, parseResult.Products[0].Categories)
	})

	t.Run("returns an error if there are more than 1 products", func(t *testing.T) {
		parseResult := SearchPage{Products: []Product{
			{"Apfelsaft", rewe.Categories{"saft"}},
			{"Apfelsaft Naturtrüb", rewe.Categories{"saft"}},
		}}
		fetcher := mockedCategoriesFetcher(t, parseResult)

		_, err := fetcher.Fetch("Apfelsaft")
		assert.ErrorType(t, err, &ErrFuzzyResult{})
	})
}

func mockedCategoriesFetcher(t *testing.T, parseResult SearchPage) CategoriesFetcher {
	getSearchPageResult := strings.NewReader("Page Content")
	mockClient := mockReweClient{t, map[string]io.Reader{
		"Apfelsaft": getSearchPageResult,
	}}

	mockParser := mockSearchPageParser{t, map[io.Reader]SearchPage{
		getSearchPageResult: parseResult,
	}}

	return CategoriesFetcher{mockClient, mockParser}
}
