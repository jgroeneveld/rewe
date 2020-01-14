package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestCategoriesCommand(t *testing.T) {
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

	assert.Equal(t, output.String(), `[
  "/c/frische-kuehlung",
  "/c/frische-kuehlung-eier-fett-molkereiprodukte",
  "/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett",
  "/c/frische-kuehlung-eier-fett-molkereiprodukte-margarine-butter-fett-butter"
]
`)
}

func newFixtureServer(t *testing.T, query string, fixture string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.String(), query) {
			t.Fatalf("unexpected request %q", r.URL.String())
		}

		f, err := os.Open("../../testdata/" + fixture)
		defer f.Close()
		assert.NilError(t, err)

		_, err = io.Copy(w, f)
		assert.NilError(t, err)
	}))
}
