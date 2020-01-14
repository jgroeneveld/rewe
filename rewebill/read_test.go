package rewebill

import (
	"gotest.tools/assert"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("../fixtures/rechnung.pdf")
	assert.NilError(t, err)
	defer f.Close()

	bill, err := Read(f)

	assert.NilError(t, err)
	assert.Equal(t, 50, len(bill.Positions))

	assert.Equal(t, Position{
		Text:   "REWE Beste Wahl Alaska-Seelachsfilets 400g",
		Amount: 1,
		Price:  Cents(299),
		Sum:    Cents(299),
		Tax:    "B",
	}, bill.Positions[0])
}