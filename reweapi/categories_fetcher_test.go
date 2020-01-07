package reweapi

import (
	"gotest.tools/assert"
	"io"
	"rewe"
	"strings"
	"testing"
)

func TestCategoriesFetcher(t *testing.T) {
	getSearchPageResult := strings.NewReader("Page Content")
	mockClient := mockReweClient{t, map[string]io.Reader{
		"Apfelsaft": getSearchPageResult,
	}}

	parseResult := SearchPage{Products: []Product{{"Apfelsaft", rewe.Categories{"saft"}}}}
	mockParser := mockSearchPageParser{t, map[io.Reader]SearchPage{
		getSearchPageResult: parseResult,
	}}

	fetcher := CategoriesFetcher{
		ReweClient:       mockClient,
		SearchPageParser: mockParser,
	}

	t.Run("returns the parse result", func(t *testing.T) {
		categories, err := fetcher.Fetch("Apfelsaft")
		assert.NilError(t, err)

		assert.DeepEqual(t, categories, parseResult.Products[0].Categories)
	})
}
