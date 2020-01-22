package main

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestCLI_Categories(t *testing.T) {
	server := newFixtureServer(t, "Butter", "search_butter.html")
	defer server.Close()

	output := bytes.NewBuffer(nil)
	app := NewApp(output)

	err := app.Run([]string{
		"rewe",
		"--base-url", server.URL,
		"categories",
		"--product", "Butter",
		"--json",
	})
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
