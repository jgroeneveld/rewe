package main

import (
	"encoding/json"
	"io"
	"rewe"

	"github.com/pkg/errors"
)

func writeFullProductInfos(w io.Writer, infos []rewe.FullProductInfo) error {
	return PrettyJSONWriter{}.Write(w, infos)
}

func writeCategoryInfo(w io.Writer, info rewe.CategoryInfo) error {
	return PrettyJSONWriter{}.Write(w, info)
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
