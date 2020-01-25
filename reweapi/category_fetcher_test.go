package reweapi

import (
	"io"

	"rewe"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestCategoryFetcher(t *testing.T) {
	t.Run("returns the parse result", func(t *testing.T) {
		saft := rewe.CategoryInfo{Categories: []string{"saft"}}
		parseResult := SearchPage{[]rewe.CategoryInfo{saft}}
		fetcher := mockedCategoryFetcher(t, parseResult)

		categories, err := fetcher.Fetch("Apfelsaft")
		assert.NilError(t, err)

		assert.DeepEqual(t, categories, parseResult.Products[0])
	})

	t.Run("returns first product if there are more than 1 products", func(t *testing.T) {
		parseResult := SearchPage{Products: []rewe.CategoryInfo{
			{Product: "Apfelsaft", Categories: []string{"saft"}},
			{Product: "Apfelsaft Naturtrüb", Categories: []string{"saft"}},
		}}
		fetcher := mockedCategoryFetcher(t, parseResult)

		categories, err := fetcher.Fetch("Apfelsaft")
		assert.NilError(t, err)

		assert.DeepEqual(t, categories, parseResult.Products[0])
	})

	t.Run("returns error if there are no products", func(t *testing.T) {
		parseResult := SearchPage{Products: []rewe.CategoryInfo{}}
		fetcher := mockedCategoryFetcher(t, parseResult)

		_, err := fetcher.Fetch("Apfelsaft")
		assert.Equal(t, err, ErrNoProductsFound)
	})
}

func mockedCategoryFetcher(t *testing.T, parseResult SearchPage) CategoryFetcher {
	getSearchPageResult := strings.NewReader("Page Content")
	mockClient := mockReweClient{t, map[string]io.Reader{
		"Apfelsaft": getSearchPageResult,
	}}

	mockParser := mockSearchPageParser{t, map[io.Reader]SearchPage{
		getSearchPageResult: parseResult,
	}}

	return CategoryFetcher{mockClient, mockParser}
}
