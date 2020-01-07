package reweapi

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

const defaultBaseUrl = "https://shop.rewe.de"

type ReweClientImpl struct {
	BaseUrl string
}

func (r ReweClientImpl) GetSearchPage(productName string) (io.Reader, error) {
	response, err := http.Get(r.searchUrl(productName))
	defer response.Body.Close()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bytes.NewReader(all), nil
}

func (r ReweClientImpl) searchUrl(query string) string {
	query = url2.QueryEscape(query)

	url := fmt.Sprintf("%s/productList?search=%s", r.getBaseUrl(), query)
	return url
}

func (r ReweClientImpl) getBaseUrl() string {
	if r.BaseUrl == "" {
		return defaultBaseUrl
	}

	return r.BaseUrl
}
