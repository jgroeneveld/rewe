package reweapi

import (
	"io"
	"testing"
)

type mockReweClient struct {
	t                  *testing.T
	GetSearchPageStubs map[string]io.Reader
}

func (m mockReweClient) GetSearchPage(productName string) (io.Reader, error) {
	m.t.Helper()

	r, ok := m.GetSearchPageStubs[productName]
	if !ok {
		m.t.Fatalf("Unsatisfied mock call %s", productName)
	}

	return r, nil
}

type mockSearchPageParser struct {
	t          *testing.T
	ParseStubs map[io.Reader]SearchPage
}

func (m mockSearchPageParser) Parse(r io.Reader) (SearchPage, error) {
	m.t.Helper()

	result, ok := m.ParseStubs[r]
	if !ok {
		m.t.Fatalf("Unsatisfied mock call %v", r)
	}

	return result, nil
}
