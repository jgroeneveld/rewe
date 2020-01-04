package main

import (
	"fmt"
	"rewe/reweapi"
)

func main() {
	fetcher := reweapi.NewCategoriesFetcher()

	categories, err := fetcher.Fetch("REWE Bio Apfelsaft naturtrüb 1l")
	if err != nil {
		panic(err.Error())
	}

	for _, c := range categories {
		fmt.Printf("%q\n", c)
	}
}

