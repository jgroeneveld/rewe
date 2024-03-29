package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}

type captureFixtureServer struct {
	*httptest.Server
	Requests []string
}

func newQueryCapturingFixtureServer(t *testing.T, fixture string) *captureFixtureServer {
	fixtureServer := &captureFixtureServer{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fixtureServer.Requests = append(fixtureServer.Requests, r.URL.String())

		f, err := os.Open("../../testdata/" + fixture)
		defer f.Close()
		assert.NilError(t, err)

		_, err = io.Copy(w, f)
		assert.NilError(t, err)
	}))

	fixtureServer.Server = server

	return fixtureServer
}

func newStrictFixtureServer(t *testing.T, query string, fixture string) *httptest.Server {
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
