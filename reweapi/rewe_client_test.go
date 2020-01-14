package reweapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestReweClientImpl(t *testing.T) {
	t.Run("sends the correct request and returns the body", func(t *testing.T) {
		gotRequestURL := ""

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotRequestURL = r.URL.String()
			fmt.Fprint(w, "the body")
		}))
		defer server.Close()

		reweClient := &ReweClientImpl{BaseURL: server.URL}

		r, err := reweClient.GetSearchPage("REWE Bio Apfelsaft naturtrüb 1l")
		assert.NilError(t, err)
		assert.Equal(
			t,
			"/productList?search=REWE+Bio+Apfelsaft+naturtr%C3%BCb+1l",
			gotRequestURL,
		)

		all, err := ioutil.ReadAll(r)
		assert.NilError(t, err)

		assert.Equal(
			t,
			"the body",
			string(all),
		)
	})
}
