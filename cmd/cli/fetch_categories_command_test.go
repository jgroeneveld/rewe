package main

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestFetchCategoriesCommand(t *testing.T) {
	server := newFixtureServer(t, "Butter", "search_butter.html")
	defer server.Close()

	output := bytes.NewBuffer(nil)

	cmd := rootCommand(output)
	cmd.SetArgs([]string{
		"--base-url", server.URL,
		"fetch-categories",
		"Butter",
		"--json",
	})

	err := cmd.Execute()
	assert.NilError(t, err)

	assert.Equal(t, output.String(), `{
  "product": "Landliebe Butter 250g",
  "categories": [
    "/c/frische-kuehlung",
    "/c/frische-kuehlung-eier-fett-molkereiprodukte",
    "/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett",
    "/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett-butter"
  ]
}
`)
}
