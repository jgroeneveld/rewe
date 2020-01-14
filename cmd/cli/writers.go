package main

import (
	"encoding/json"
	"fmt"
	"io"
	"rewe"

	"github.com/pkg/errors"
)

func writeCategories(w io.Writer, categories rewe.Categories, useJSON bool) error {
	if useJSON {
		return PrettyJSONWriter{}.Write(w, categories)
	}

	return SimpleCategoriesWriter{}.WriteCategories(w, categories)
}

// SimpleCategoriesWriter writes categories line by line as quoted strings
type SimpleCategoriesWriter struct {
}

// WriteCategories writes categories line by line as quoted strings
func (w SimpleCategoriesWriter) WriteCategories(writer io.Writer, categories rewe.Categories) error {
	for _, c := range categories {
		fmt.Fprintf(writer, "%q\n", c)
	}

	return nil
}

// PrettyJSONWriter writes the given data as pretty JSON
type PrettyJSONWriter struct {
}

// Write writes the given data as pretty JSON
func (w PrettyJSONWriter) Write(writer io.Writer, data interface{}) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(data)
	if err != nil {
		return errors.Wrap(err, "can not write json")
	}

	return nil
}
