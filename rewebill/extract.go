package rewebill

import (
	"regexp"
	"rewe"
	"rewe/util/check"
	"strconv"
	"strings"
)

func Extract(pdf Pdf) (rewe.Bill, error) {
	var positions []rewe.Position

	for _, line := range pdf.AllLines() {
		position, ok := isPositionLine(line)
		if ok {
			positions = append(positions, position)
		}
	}

	return rewe.Bill{Positions: positions}, nil
}

func isPositionLine(line string) (rewe.Position, bool) {
	subMatches := positionLineRegex.FindStringSubmatch(line)
	if len(subMatches) == 0 {
		return rewe.Position{}, false
	}

	check.Equal(len(subMatches), 6, "number of matches does not match regex")

	return rewe.Position{
		Text:   subMatches[1],
		Amount: asInt(subMatches[2]),
		Price:  asCents(subMatches[3]),
		Sum:    asCents(subMatches[4]),
		Tax:    subMatches[5],
	}, true
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
