
package Controller

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Square string
type SquarePossibilities map[Square]string

var digits, rows, cols = "123456789", "ABCDEFGHI", "123456789"
var squares []Square
var unitlist [][]Square
var peers map[Square][]Square
var units map[Square][][]Square
var validCharRegex *regexp.Regexp

func init() {
	initialize()
}
func initialize() {
	validCharRegex = regexp.MustCompile(`[0-9]*\,*`)
	squares = Cross(rows, cols)
	unitlist = BuildUnitList(rows, cols, []string{"ABC", "DEF", "GHI"}, []string{"123", "456", "789"})
	units = map[Square][][]Square{}
	for _, square := range squares {
	Square:
		for _, v := range unitlist {
			for _, unitsquare := range v {
				if square == unitsquare {
					units[square] = append(units[square], v)
					continue Square
				}
			}
		}
	}
	peers = map[Square][]Square{}
	for square, unitlist := range units {
		for _, unit := range unitlist {
		NextUnitSquare:
			for _, unitsquare := range unit {
				if unitsquare != square {
					for _, ps := range peers[square] {
						if ps == unitsquare {
							continue NextUnitSquare
						}
					}

					peers[square] = append(peers[square], unitsquare)
				}
			}
		}
	}
}
func Cross(A string, B string) []Square {
	var out []Square
	for _, i := range A {
		for _, j := range B {
			out = append(out, Square(i)+Square(j))
		}
	}

	return out
}
func BuildUnitList(rows string, cols string, rowBlocks []string, rowCols []string) [][]Square {
	out := [][]Square{}
	for _, c := range cols {
		out = append(out, Cross(rows, string(c)))
	}

	for _, r := range rows {
		out = append(out, Cross(string(r), cols))
	}

	for _, rs := range rowBlocks {
		for _, cs := range rowCols {
			out = append(out, Cross(rs, cs))
		}
	}

	return out
}
func ParseGrid(Grid string) (SquarePossibilities, error) {
	values := SquarePossibilities{}
	for _, square := range squares {
		values[square] = digits
	}

	vGrid := strings.Join(validCharRegex.FindAllString(Grid, -1), "")
	gridMap, err := gridSquarePossibilities(vGrid)
	if err != nil {
		return nil, err
	}
	for s, d := range gridMap {
		for _, xd := range digits {
			if d == string(xd) {
				values, err = assign(values, s, d)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return values, nil
}
func gridSquarePossibilities(vGrid string) (SquarePossibilities, error) {
	out := SquarePossibilities{}
	if len(vGrid) != 81 {
		return SquarePossibilities{}, errors.New("vGrid does not contain 81 valid digits.")
	}
	for i := 0; i < 81; i++ {
		if string(vGrid[i]) == "0" {
			out[squares[i]] = "."
		} else {
			out[squares[i]] = string(vGrid[i])
		}
	}

	return out, nil
}
func assign(values SquarePossibilities, s Square, d string) (SquarePossibilities, error) {
	if len(values[s]) < 1 {
		return nil, errors.New(fmt.Sprintf("values[%s] has no permissible digits left; contradiction.", s))
	}

	otherSquarePossibilities := ``
	for _, v := range values[s] {
		if string(v) == d {
			continue
		}

		otherSquarePossibilities = otherSquarePossibilities + string(v)
	}

	if len(otherSquarePossibilities) > 0 {
		for _, d2 := range otherSquarePossibilities {
			if _, err := eliminate(values, s, string(d2)); err != nil {
				return nil, err
			}
		}
	}

	return values, nil
}
func eliminate(values SquarePossibilities, s Square, d string) (SquarePossibilities, error) {
	err := error(nil)

	dInSquarePossibilitiesS := false
	for _, val := range values[s] {
		if string(val) == d {
			dInSquarePossibilitiesS = true
			break
		}
	}

	if !dInSquarePossibilitiesS {
		return values, nil
	}

	values[s] = strings.Replace(values[s], d, "", -1)

	if len(values[s]) == 0 {
		return nil, errors.New(fmt.Sprintf("Cannot eliminate %s from values[%s] because now values[%s] has no valid potential digits (valid: %s).", d, s, s, values[s]))
	} else if len(values[s]) == 1 {
		d2 := values[s]
		for _, s2 := range peers[s] {
			if _, err = eliminate(values, s2, d2); err != nil {
				return nil, err
			}
		}
	}

	for _, u := range units[s] {
		dPlaces := []Square{}
		for _, s2 := range u {
			for _, d2 := range values[s2] {
				if d == string(d2) {
					dPlaces = append(dPlaces, s2)
					break
				}
			}
		}

		if len(dPlaces) == 0 {
			return values, errors.New(fmt.Sprintf("There is no place in unit %s to put %s.", u, d))
		} else if len(dPlaces) == 1 {
			_, err = assign(values, dPlaces[0], d)
			if err != nil {
				return nil, err
			}
		}
	}

	return values, nil
}

func Solve(Grid string) (SquarePossibilities, error) {
	values, err := ParseGrid(Grid)
	if err != nil {
		return values, err
	}

	return search(values, nil)
}
func search(values SquarePossibilities, err error) (SquarePossibilities, error) {
	if err != nil {
		return values, err
	}

	solved := true
	for _, s := range squares {
		if len(values[s]) != 1 {
			solved = false
		}
	}
	if solved {
		return values, nil
	}
	min, sq := len(digits)+1, Square("")
	for _, s := range squares {
		if len(values[s]) < min && len(values[s]) > 1 {
			min = len(values[s])
			sq = s

			if min == 2 {
				break
			}
		}
	}

	for _, d := range values[sq] {

		vCloned, err := assign(cloneSquarePossibilities(values), sq, string(d))

		if err != nil {
			continue
		}
		vCloned, err = search(vCloned, err)
		if err == nil {
			return vCloned, nil
		}
	}

	return nil, errors.New("Your depth-first search failed on this branch.")
}

func cloneSquarePossibilities(values SquarePossibilities) SquarePossibilities {
	cpySquarePossibilities := make(SquarePossibilities, len(values))
	for k, v := range values {
		cpySquarePossibilities[k] = v
	}

	return cpySquarePossibilities
}
func Display(values SquarePossibilities) string {
	mLen := 0
	for _, s := range squares {
		if len(values[s]) > mLen {
			mLen = len(values[s])
		}
	}
	width := mLen + 1

	line := []byte{}
	for loop := 0; loop < 3; loop++ {
		for i := 0; i < width*3; i++ {
			line = append(line, []byte("-")...)
		}
		if loop < 2 {
			line = append(line, []byte("+")...)
		}
	}
	line = append(line, []byte("\n")...)

	out := []byte{}
	for rn, r := range rows {
		for cn, c := range cols {
			s := Square(fmt.Sprintf("%s%s", string(r), string(c)))
			out = append(out, []byte(strWid(values[s], width))...)

			if cn == 2 || cn == 5 {
				out = append(out, []byte(`|`)...)
			}
		}
		out = append(out, []byte("\n")...)
		if rn == 2 || rn == 5 {
			out = append(out, line...)
		}
	}
	return string(out)
}
func strWid(s string, width int) string {
	if len(s) >= width {
		return s
	}

	for i := len(s); i < width; i++ {
		s += " "
	}
	return s
}
