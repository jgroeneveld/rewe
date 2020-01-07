package reweapi

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"rewe"
	"strings"
)

type SearchPage struct {
	Products []Product
}

type Product struct {
	Name       string
	Categories rewe.Categories
}

type SearchPageParserImpl struct{}

func (p SearchPageParserImpl) Parse(r io.Reader) (SearchPage, error) {
	lines, err := readLines(r)
	if err != nil {
		return SearchPage{}, err
	}

	jsonString, err := extractJsonString(lines)
	if err != nil {
		return SearchPage{}, err
	}

	parsedJson, err := parseJson(jsonString)
	if err != nil {
		return SearchPage{}, err
	}

	return SearchPage{
		Products: parsedJson.Products(),
	}, nil
}

func readLines(r io.Reader) ([]string, error) {
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "Cant read")
	}

	return strings.Split(string(all), "\n"), nil
}

func extractJsonString(lines []string) (string, error) {
	dataLine := ""

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

func parseJson(jsonString string) (*searchPageJson, error) {
	var parsedJson *searchPageJson
	err := json.NewDecoder(strings.NewReader(jsonString)).Decode(&parsedJson)
	if err != nil {
		return nil, errors.Wrap(err, "can not decode json")
	}

	return parsedJson, nil
}

type searchPageJson struct {
	Embedded struct {
		Products []*productJson `json:"products"`
	} `json:"_embedded"`
}

func (j searchPageJson) Products() []Product {
	products := []Product{}

	for _, p := range j.Embedded.Products {
		products = append(products, p.AsModel())
	}
	return products
}

type productJson struct {
	ProductName string `json:"productName"`
	Embedded    struct {
		Categories []*categoryJson
	} `json:"_embedded"`
}

func (j productJson) AsModel() Product {
	return Product{
		Name:       j.ProductName,
		Categories: j.CategoriesAsModels(),
	}
}

func (j productJson) CategoriesAsModels() rewe.Categories {
	categories := rewe.Categories{}

	for _, c := range j.Embedded.Categories {
		categories = append(categories, c.Links.Products.Href)
	}

	return categories
}

type categoryJson struct {
	Id    string `json:"id"`
	Links struct {
		Products struct {
			Href string `json:"href"`
		} `json:"products"`
	} `json:"_links"`
}
