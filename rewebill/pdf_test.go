package rewebill

import (
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestReadPdf(t *testing.T) {
	f, err := os.Open("../testdata/rechnung.pdf")
	assert.NilError(t, err)
	defer f.Close()

	pdf, err := ReadPdf(f)
	assert.NilError(t, err)

	assert.Equal(t, 3, len(pdf.TextPages))
	assert.Assert(t, strings.Contains(pdf.TextPages[0], "Iglo Dill 50g"))
	assert.Assert(t, strings.Contains(pdf.TextPages[1], "Pepsi Max Zero 6x1,5l"))
	assert.Assert(t, strings.Contains(pdf.TextPages[2], "Vielen Dank für Ihren Einkauf bei REWE"))
}
