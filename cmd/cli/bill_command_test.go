package main

import (
	"bytes"
	"testing"

	"github.com/jgroeneveld/schema"

	"gotest.tools/assert"
)

func TestBillCommand(t *testing.T) {
	captureServer := newQueryCapturingFixtureServer(t, "search_butter.html")
	defer captureServer.Close()

	output := bytes.NewBuffer(nil)

	cmd := rootCommand(output)
	cmd.SetArgs([]string{
		"bill",
		"../../testdata/rechnung.pdf",
		"--base-url", captureServer.URL,
	})

	err := cmd.Execute()
	assert.NilError(t, err)

	err = schema.MatchJSON(
		schema.Map{
			"order_date": "08.01.2020",
			"augmented_positions": schema.ArrayEach(
				schema.Map{
					"text":       schema.IsString,
					"categories": schema.ArrayEach(schema.IsString),
					"amount":     schema.IsInteger,
					"price":      schema.IsInteger,
					"sum":        schema.IsInteger,
					"tax":        schema.IsString,
				},
			),
		},
		output,
	)
	assert.NilError(t, err)

	assert.DeepEqual(t, captureServer.Requests, []string{
		"/productList?search=REWE+Beste+Wahl+Alaska-Seelachsfilets+400g",
		"/productList?search=Iglo+Dill+50g",
		"/productList?search=ja%21+Kartoffelchips+mit+Salz+200g",
		"/productList?search=ja%21+Kartoffel-Chips+mit+Paprika+200g",
		"/productList?search=ja%21+Erdn%C3%BCsse+ger%C3%B6stet+%26+gesalzen+200g",
		"/productList?search=REWE+Regional+Apfel+rot+1kg",
		"/productList?search=REWE+Bio+Rispentomaten+500g",
		"/productList?search=REWE+Bio+Gurke",
		"/productList?search=REWE+Beste+Wahl+Romana+Salatherzen+3",
		"/productList?search=Avocado+essreif+2+St%C3%BCck+in+Schale",
		"/productList?search=Wawi+Blockschokolade+Zartbitter+200g",
		"/productList?search=REWE+Bio+Maiswaffeln+115g",
		"/productList?search=REWE+Beste+Wahl+WeizenTortillas+432g",
		"/productList?search=REWE+Beste+Wahl+Weizenmehl+Type+405+1kg",
		"/productList?search=REWE+Beste+Wahl+Thunfisch-Filets+in",
		"/productList?search=REWE+Beste+Wahl+st%C3%BCckige+Tomaten+425ml",
		"/productList?search=REWE+Beste+Wahl+Kichererbsen+265g",
		"/productList?search=Nissin+Nudelsuppe+Huhn+100g",
		"/productList?search=Nissin+Cup+Noodles+Huhn+63g",
		"/productList?search=Maggi+Magic+Asia+Instant+Nudel+Snack+Huhn",
		"/productList?search=Leicht+%26+Cross+Knusperbrot+Weizen+125g",
		"/productList?search=K%C3%BChne+Wei%C3%9Fwein-Essig+500ml",
		"/productList?search=K%C3%BChne+K%C3%BCrbis+200g",
		"/productList?search=ja%21+Super+Sweet+Gem%C3%BCsemais+425ml",
		"/productList?search=ja%21+Basmati+Reis+1kg",
		"/productList?search=Bamboo+Garden+Mango-Chutney+mild+230g",
		"/productList?search=Vernel+Weichsp%C3%BCler+Frischer+Morgen+1l%2C+33WL",
		"/productList?search=Swirl+Staubsaugerbeutel+S+67+MicroPor+Plus+4",
		"/productList?search=Dalli+Colorwaschmittel+1%2C1l%2C+20+WL",
		"/productList?search=REWE+Beste+Wahl+Zitronensaft+0%2C75l",
		"/productList?search=REWE+Beste+Wahl+Limonata+Zitrone+0%2C33l",
		"/productList?search=REWE+Beste+Wahl+Aranciata+Orangen+0%2C33l",
		"/productList?search=Pepsi+Max+Zero+6x1%2C5l",
		"/productList?search=Wiltmann+Bio-Gefl%C3%BCgel-Lyoner+80g",
		"/productList?search=Seraphos+Joghurt+griechische+Art+4x100g",
		"/productList?search=REWE+Bio+Fettarme+H-Milch+1%2C5%25+1l",
		"/productList?search=REWE+Beste+Wahl+H%C3%A4hnchenfleisch+in+Aspik",
		"/productList?search=REWE+Beste+Wahl+H-Sahne+zum+Kochen+15%25",
		"/productList?search=REWE+Beste+Wahl+Gnocchi+400g",
		"/productList?search=Landliebe+Butter+250g",
		"/productList?search=ja%21+Butterk%C3%A4se+400g",
		"/productList?search=Du+darfst+Apfel-Zwiebel-Leberwurst+100g",
		"/productList?search=REWE+Beste+Wahl+Windeln+Maxi+42+St%C3%BCck",
	})
}
