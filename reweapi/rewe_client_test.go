package reweapi

import (
	"fmt"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReweClientImpl(t *testing.T) {
	t.Run("sends the correct request and returns the body", func(t *testing.T) {
		gotRequestUrl := ""

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotRequestUrl = r.URL.String()
			fmt.Fprint(w, "the body")
		}))
		defer server.Close()

		reweClient := &ReweClientImpl{BaseUrl: server.URL}

		r, err := reweClient.GetSearchPage("REWE Bio Apfelsaft naturtrüb 1l")
		assert.NilError(t, err)
		assert.Equal(
			t,
			"/productList?search=REWE+Bio+Apfelsaft+naturtr%C3%BCb+1l",
			gotRequestUrl,
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
