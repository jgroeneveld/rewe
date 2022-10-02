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
		"../../testdata/rechnung-blackend.pdf",
	})

	err := cmd.Execute()
	assert.NilError(t, err)

	err = schema.MatchJSON(
		schema.Map{
			"order_date": "08.01.2020",
			"positions": schema.ArrayIncluding(
				schema.Map{
					"text":   "REWE Beste Wahl Alaska-Seelachsfilets 400g",
					"amount": 1,
					"price":  299,
					"sum":    299,
					"tax":    "B",
				},
				schema.Map{
					"text":   "REWE Beste Wahl Weizenmehl Type 405 1kg",
					"amount": 1,
					"price":  59,
					"sum":    59,
					"tax":    "B",
				},
			),
		},
		output,
	)
	assert.NilError(t, err)
}
