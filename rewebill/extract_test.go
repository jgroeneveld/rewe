package rewebill

import (
	"rewe"
	"testing"

	"gotest.tools/assert"
)

func TestExtract(t *testing.T) {
	bill, err := Extract(pdfFixture)
	assert.NilError(t, err)

	expected := rewe.Bill{
		OrderDate: "02.01.2020",
		Positions: []rewe.Position{
			{
				Text:   "REWE Beste Wahl Alaska-Seelachsfilets 400g",
				Amount: 1,
				Price:  rewe.Cents(299),
				Sum:    rewe.Cents(299),
				Tax:    "B",
			},
			{
				Text:   "REWE Bio Maiswaffeln 115g",
				Amount: 1,
				Price:  rewe.Cents(79),
				Sum:    rewe.Cents(79),
				Tax:    "B",
			},
			{
				Text:   "Wiltmann Bio-Geflügel-Lyoner 80g",
				Amount: 2,
				Price:  rewe.Cents(169),
				Sum:    rewe.Cents(338),
				Tax:    "B",
			},
		},
	}
	assert.DeepEqual(t, expected, bill)
}

var pdfFixture = Pdf{TextPages: []string{`
Rechnung OL20064800054179
Bestelldatum: 02.01.2020
Bestellnummer: B-SQL-2ZV-SYB
Bezeichnung Menge Einzelpreis Summe Pos. MwSt.
REWE Beste Wahl Alaska-Seelachsfilets 400g 1 2,99 € 2,99 € B
REWE Bio Maiswaffeln 115g 1 0,79 € 0,79 € B
Wiltmann Bio-Geflügel-Lyoner 80g 2 1,69 € 3,38 € B
`, `
Servicegebühr Lieferung 1 0,99 € 0,99 € A
Servicegebühr Lieferung 1 1,91 € 1,91 € B
Leergut -1 1,50 € -1,50 € A
PFAND 1,50 1 1,50 € 1,50 € A
Summe: 92,65 €
Summe nach Steuersatz
Steuersatz Nettobetrag MwSt. Betrag Gesamtbetrag
A = 19% 26,25 € 4,99 € 31,24 €
B = 7% 57,39 € 4,02 € 61,41 €
Summe: 92,65 €
`, `
Vielen Dank für Ihren Einkauf bei REWE
`}}
