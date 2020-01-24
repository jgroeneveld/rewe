package main

import (
	"bytes"
	"testing"

	"github.com/jgroeneveld/schema"

	"gotest.tools/assert"
)

func TestReadBillCommand(t *testing.T) {
	output := bytes.NewBuffer(nil)

	cmd := rootCommand(output)
	cmd.SetArgs([]string{
		"read-bill",
		"../../testdata/rechnung.pdf",
	})

	err := cmd.Execute()
	assert.NilError(t, err)

	err = schema.MatchJSON(
		schema.Map{
			"Positions": schema.ArrayIncluding(
				schema.Map{
					"Text":   "REWE Beste Wahl Alaska-Seelachsfilets 400g",
					"Amount": 1,
					"Price":  299,
					"Sum":    299,
					"Tax":    "B",
				},
				schema.Map{
					"Text":   "REWE Beste Wahl Weizenmehl Type 405 1kg",
					"Amount": 1,
					"Price":  59,
					"Sum":    59,
					"Tax":    "B",
				},
			),
		},
		output,
	)
	assert.NilError(t, err)
}
