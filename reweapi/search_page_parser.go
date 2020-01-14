package reweapi

import (
	"bufio"
	"encoding/json"
	"io"
	"rewe"
	"strings"

	"github.com/pkg/errors"
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
	jsonString, err := extractJSONString(r)
	if err != nil {
		return SearchPage{}, err
	}

	parsedJSON, err := parseJSON(jsonString)
	if err != nil {
		return SearchPage{}, err
	}

	return SearchPage{
		Products: parsedJSON.Products(),
	}, nil
}

func extractJSONString(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	dataLine := ""

	for scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text())
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

func (j searchPageJSON) Products() []Product {
	products := []Product{}

	for _, p := range j.Embedded.Products {
		products = append(products, p.AsModel())
	}
	return products
}

type productJSON struct {
	ProductName string `json:"productName"`
	Embedded    struct {
		Categories []*categoryJSON
	} `json:"_embedded"`
}

func (j productJSON) AsModel() Product {
	return Product{
		Name:       j.ProductName,
		Categories: j.CategoriesAsModels(),
	}
}

func (j productJSON) CategoriesAsModels() rewe.Categories {
	categories := rewe.Categories{}

	for _, c := range j.Embedded.Categories {
		categories = append(categories, c.Links.Products.Href)
	}

	return categories
}

type categoryJSON struct {
	ID    string `json:"id"`
	Links struct {
		Products struct {
			Href string `json:"href"`
		} `json:"products"`
	} `json:"_links"`
}
