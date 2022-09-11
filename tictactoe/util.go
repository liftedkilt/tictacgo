package tictactoe

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Marker string

const (
	X Marker = "X"
	O Marker = "O"
	B Marker = " "
)

type Board [3][3]Marker

type Position struct {
	Col int
	Row int
}

func New() Board {
	row := [3]Marker{B, B, B}

	var b Board

	for i := 0; i < 3; i++ {
		b[i] = row
	}

	return b
}

func (b Board) String() string {
	var sb strings.Builder

	sb.WriteString("   A   B   C\n")

	for i, row := range b {
		s := fmt.Sprintf("%d  %s | %s | %s\n", i+1, row[0], row[1], row[2])

		if i == 1 {
			sb.WriteString("   ----------\n")
		}

		sb.WriteString(s)

		if i == 1 {
			sb.WriteString("   ----------\n")
		}
	}

	return sb.String()
}

func (b Board) Mark(p Position, m Marker) (Board, error) {
	if b[p.Row][p.Col] != B {
		return b, errors.New("space occupied")
	}

	b[p.Row][p.Col] = m

	return b, nil
}

func ParsePosition(s string) (Position, error) {
	r := regexp.MustCompile(`(?P<col>[a-cA-C])(?P<row>[1-3])`)

	matches := r.FindStringSubmatch(s)

	if len(matches) != 3 {
		return Position{}, errors.New("invalid coordinates")
	}

	columnString := matches[1]

	columnString = strings.ToUpper(columnString)

	var col int

	switch columnString {
	case "A":
		col = 0
	case "B":
		col = 1
	case "C":
		col = 2
	default:
		panic("unreachable")
	}

	row, err := strconv.Atoi(matches[2])

	row = row - 1

	if err != nil {
		panic("unreachable")
	}

	p := Position{
		Col: col,
		Row: row,
	}

	return p, nil
}

func (b Board) IsEndState() bool {
	// Check for Row victory
	for _, row := range b {
		if row[0] == B {
			continue
		}

		if row[0] == row[1] && row[1] == row[2] {
			return true
		}
	}

	// Check for Column victory
	for i := 0; i < 3; i++ {
		if b[0][i] == B {
			continue
		}

		if b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return true
		}
	}

	// Check Diagonal victory
	if b[1][1] == B {
		return false
	}

	if (b[0][0] == b[1][1] && b[1][1] == b[2][2]) ||
		(b[0][2] == b[1][1] && b[1][1] == b[2][0]) {
		return true
	}

	return false
}
