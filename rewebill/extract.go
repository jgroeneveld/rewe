package rewebill

import (
	"regexp"
	"rewe"
	"rewe/util/check"
	"strconv"
	"strings"
)

func Extract(pdf Pdf) (rewe.Bill, error) {
	var orderDate string
	var positions []rewe.Position

	for _, line := range pdf.AllLines() {
		date, ok := isOrderDateLine(line)
		if ok {
			orderDate = date
		}

		position, ok := isPositionLine(line)
		if ok {
			positions = append(positions, position)
		}
	}

	return rewe.Bill{
		OrderDate: orderDate,
		Positions: positions,
	}, nil
}

func isOrderDateLine(line string) (string, bool) {
	if !strings.Contains(line, "Bestelldatum") {
		return "", false
	}

	parts := strings.Split(line, " ")
	check.Equal(len(parts), 3, "number of parts in orderDate line is not correct")

	return parts[2], true
}

func isPositionLine(line string) (rewe.Position, bool) {
	subMatches := positionLineRegex.FindStringSubmatch(line)
	if len(subMatches) == 0 {
		return rewe.Position{}, false
	}

	check.Equal(len(subMatches), 6, "number of matches does not match regex")

	if isBlacklisted(subMatches[1]) {
		return rewe.Position{}, false
	}

	return rewe.Position{
		Text:   subMatches[1],
		Amount: asInt(subMatches[2]),
		Price:  asCents(subMatches[3]),
		Sum:    asCents(subMatches[4]),
		Tax:    subMatches[5],
	}, true
}

func isBlacklisted(text string) bool {
	text = strings.ToLower(text)

	for _, blacklisted := range positionBlacklist {
		if strings.Contains(text, blacklisted) {
			return true
		}
	}

	return false
}

func asCents(s string) rewe.Cents {
	s = strings.ReplaceAll(strings.ReplaceAll(s, " €", ""), ",", "")
	return rewe.Cents(asInt(s))
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	check.Error(err)
	return i
}

var positionLineRegex = regexp.MustCompile(`(.+) (-?\d+) (-?\d+,\d+ €) (-?\d+,\d+ €) (\w)$`)
var positionBlacklist = []string{"leergut", "servicegebühr", "pfand"}
