package rewebill

import (
	"regexp"
	"rewe/util/check"
	"strconv"
	"strings"
)

type Bill struct {
	Positions []Position
}

type Position struct {
	Text   string
	Amount int
	Price  Cents
	Sum    Cents
	Tax    string
}

type Cents int

func Extract(pdf Pdf) (Bill, error) {
	var positions []Position

	for _, line := range pdf.AllLines() {
		position, ok := isPositionLine(line)
		if ok {
			positions = append(positions, position)
		}
	}

	return Bill{Positions: positions}, nil
}

func isPositionLine(line string) (Position, bool) {
	subMatches := positionLineRegex.FindStringSubmatch(line)
	if len(subMatches) == 0 {
		return Position{}, false
	}

	check.Equal(len(subMatches), 6, "number of matches does not match regex")

	return Position{
		Text:   subMatches[1],
		Amount: asInt(subMatches[2]),
		Price:  asCents(subMatches[3]),
		Sum:    asCents(subMatches[4]),
		Tax:    subMatches[5],
	}, true
}

func asCents(s string) Cents {
	s = strings.ReplaceAll(strings.ReplaceAll(s, " €", ""), ",", "")
	return Cents(asInt(s))
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	check.Error(err)
	return i
}

var positionLineRegex = regexp.MustCompile(`(.+) (-?\d+) (-?\d+,\d+ €) (-?\d+,\d+ €) (\w)$`)
