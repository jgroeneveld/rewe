package reweapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"rewe"
	"strings"

	"github.com/pkg/errors"
)

type SearchPage struct {
	Products []rewe.CategoryInfo
}

type SearchPageParserImpl struct{}

func (p SearchPageParserImpl) Parse(r io.Reader) (SearchPage, error) {
	jsonString, err := extractJSONString(r)
	if err != nil {
		return SearchPage{}, err
	}

	parsedJSON, err := parseJSON(jsonString)
	if err != nil {
		return SearchPage{}, err
	}

	return SearchPage{
		Products: parsedJSON.AsModel(),
	}, nil
}

func extractJSONString(r io.Reader) (string, error) {
	dataLine := ""

	all, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(all), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "renderClientSide") {
			dataLine = trimmed
			break
		}
	}

	if dataLine == "" {
		return "", errors.New("Can not find dataline")
	}

	dataLine = strings.TrimPrefix(dataLine, "renderClientSide(")
	dataLine = strings.TrimSuffix(dataLine, ");")

	return dataLine, nil
}

func parseJSON(jsonString string) (*searchPageJSON, error) {
	var parsedJSON *searchPageJSON
	err := json.NewDecoder(strings.NewReader(jsonString)).Decode(&parsedJSON)
	if err != nil {
		return nil, errors.Wrap(err, "can not decode json")
	}

	return parsedJSON, nil
}

type searchPageJSON struct {
	Embedded struct {
		Products []*productJSON `json:"products"`
	} `json:"_embedded"`
}

func (j searchPageJSON) AsModel() []rewe.CategoryInfo {
	infos := []rewe.CategoryInfo{}

	for _, p := range j.Embedded.Products {
		infos = append(infos, p.AsModel())
	}
	return infos
}

type productJSON struct {
	ProductName string `json:"productName"`
	Embedded    struct {
		Categories []*categoryJSON
	} `json:"_embedded"`
}

func (j productJSON) AsModel() rewe.CategoryInfo {
	var categories []string

	for _, c := range j.Embedded.Categories {
		categories = append(categories, c.Links.Products.Href)
	}

	return rewe.CategoryInfo{
		Product:    j.ProductName,
		Categories: categories,
	}
}

type categoryJSON struct {
	ID    string `json:"id"`
	Links struct {
		Products struct {
			Href string `json:"href"`
		} `json:"products"`
	} `json:"_links"`
}
