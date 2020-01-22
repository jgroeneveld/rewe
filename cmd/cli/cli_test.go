package main

import (
	"gotest.tools/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
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