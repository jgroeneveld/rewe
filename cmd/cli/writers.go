package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"rewe"
)

func writeCategories(w io.Writer, categories rewe.Categories, useJson bool) error {
	var categoriesWriter CategoriesWriter

	if useJson {
		categoriesWriter = JsonCategoriesWriter{}
	} else {
		categoriesWriter = SimpleCategoriesWriter{}
	}

	return categoriesWriter.WriteCategories(w, categories)
}

type CategoriesWriter interface {
	WriteCategories(io.Writer, rewe.Categories) error
}

type SimpleCategoriesWriter struct {
}

func (w SimpleCategoriesWriter) WriteCategories(writer io.Writer, categories rewe.Categories) error {
	for _, c := range categories {
		fmt.Fprintf(writer, "%q\n", c)
	}

	return nil
}

type JsonCategoriesWriter struct {
}

func (w JsonCategoriesWriter) WriteCategories(writer io.Writer, categories rewe.Categories) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(categories)
	if err != nil {
		return errors.Wrap(err, "can not write json")
	}

	return nil
}
