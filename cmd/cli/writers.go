package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"rewe"
)

func writeCategories(w io.Writer, categories rewe.Categories, useJson bool) error {
	if useJson {
		return PrettyJsonWriter{}.Write(w, categories)
	} else {
		return SimpleCategoriesWriter{}.WriteCategories(w, categories)
	}
}

type SimpleCategoriesWriter struct {
}

func (w SimpleCategoriesWriter) WriteCategories(writer io.Writer, categories rewe.Categories) error {
	for _, c := range categories {
		fmt.Fprintf(writer, "%q\n", c)
	}

	return nil
}

type PrettyJsonWriter struct {
}

func (w PrettyJsonWriter) Write(writer io.Writer, data interface{}) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(data)
	if err != nil {
		return errors.Wrap(err, "can not write json")
	}

	return nil
}
