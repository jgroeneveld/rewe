package main

import (
	"encoding/json"
	"fmt"
	"io"
	"rewe"

	"github.com/pkg/errors"
)

func writeCategoryInfos(w io.Writer, infos []rewe.CategoryInfo, useJSON bool) error {
	if useJSON {
		return PrettyJSONWriter{}.Write(w, infos)
	}

	for _, info := range infos {
		err := SimpleCategoryInfoWriter{}.WriteCategoryInfo(w, info)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeCategoryInfo(w io.Writer, info rewe.CategoryInfo, useJSON bool) error {
	if useJSON {
		return PrettyJSONWriter{}.Write(w, info)
	}

	return SimpleCategoryInfoWriter{}.WriteCategoryInfo(w, info)
}

// SimpleCategoryInfoWriter writes categories line by line as quoted strings
type SimpleCategoryInfoWriter struct {
}

// WriteCategoryInfo writes categories line by line as quoted strings
func (w SimpleCategoryInfoWriter) WriteCategoryInfo(writer io.Writer, info rewe.CategoryInfo) error {
	for _, c := range info.Categories {
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
