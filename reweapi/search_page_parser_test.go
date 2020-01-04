package reweapi

import (
	"gotest.tools/assert"
	"os"
	"rewe"
	"testing"
)

func TestSearchPageParser(t *testing.T) {
	file, err := os.Open("../fixtures/search_butter.html")
	assert.NilError(t, err)
	defer file.Close()

	parser := SearchPageParserImpl{}

	searchPage, err := parser.Parse(file)
	assert.NilError(t, err)
	assert.DeepEqual(t, searchPage, &SearchPage{[]*Product{
		{
			Name: "Landliebe Butter 250g",
			Categories: rewe.Categories{
				"/c/frische-kuehlung",
				"/c/frische-kuehlung-eier-fett-molkereiprodukte",
				"/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett",
				"/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett-butter",
			}},
	}})
}
