package reweapi

import (
	"io"
	"os"
	"rewe"
	"testing"

	"gotest.tools/assert"
)

func TestSearchPageParser(t *testing.T) {
	parser := SearchPageParserImpl{}

	t.Run("parses products when data present", func(t *testing.T) {
		file := openFixture(t, "search_butter.html")
		defer file.Close()

		searchPage, err := parser.Parse(file)
		assert.NilError(t, err)
		assert.DeepEqual(t, searchPage, SearchPage{[]Product{
			{
				Name: "Landliebe Butter 250g",
				Categories: rewe.Categories{
					"/c/frische-kuehlung",
					"/c/frische-kuehlung-eier-fett-molkereiprodukte",
					"/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett",
					"/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett-butter",
				}},
		}})
	})

	t.Run("returns error when data is missing", func(t *testing.T) {
		file := openFixture(t, "search_missing_data.html")
		defer file.Close()

		_, err := parser.Parse(file)
		assert.ErrorContains(t, err, "Can not find dataline")
	})
}

func openFixture(t *testing.T, fileName string) io.ReadCloser {
	t.Helper()

	file, err := os.Open("../testdata/" + fileName)
	assert.NilError(t, err)

	return file
}
