package reweapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const defaultBaseURL = "https://shop.rewe.de"

type ReweClientImpl struct {
	BaseURL string
}

func (r ReweClientImpl) GetSearchPage(productName string) (io.Reader, error) {
	logger := log.WithField("Caller", "ReweClient.GetSearchPage")

	url := r.searchURL(productName)
	logger.Debugf("GET %q", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer response.Body.Close()

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bytes.NewReader(all), nil
}

func (r ReweClientImpl) searchURL(query string) string {
	query = url2.QueryEscape(query)

	url := fmt.Sprintf("%s/productList?search=%s", r.getBaseURL(), query)
	return url
}

func (r ReweClientImpl) getBaseURL() string {
	if r.BaseURL == "" {
		return defaultBaseURL
	}

	return r.BaseURL
}
