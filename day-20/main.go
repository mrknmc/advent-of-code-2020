package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tile struct {
	variants [][][]rune
}

func (tile *Tile) topBorders() [][]rune {
	result := make([][]rune, 0)
	for _, v := range tile.variants {
		result = append(result, v[0])
	}
	return result
}

func (tile *Tile) bottomBorders() [][]rune {
	result := make([][]rune, 0)
	for _, v := range tile.variants {
		result = append(result, v[len(v)-1])
	}
	return result
}

func verticalStitch(tile *Tile, other *Tile) (int, int, bool) {
	for i, topBorder := range tile.topBorders() {
		for j, bottomBorder := range other.bottomBorders() {
			// conversion to string
			if string(topBorder) == string(bottomBorder) {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func rotate(bitmap [][]rune) [][]rune {
	result := make([][]rune, len(bitmap))
	for i, row := range bitmap {
		result[i] = make([]rune, len(row))
	}

	lastRow := len(bitmap) - 1
	for i, row := range bitmap {
		for j, char := range row {
			result[j][lastRow-i] = char
		}
	}
	return result
}

func flipV(rows [][]rune) [][]rune {
	result := make([][]rune, len(rows))
	for i := range rows {
		result[i] = rows[len(rows)-1-i]
	}
	return result
}

func flipH(rows [][]rune) [][]rune {
	result := make([][]rune, len(rows))
	for i, row := range rows {
		result[i] = make([]rune, len(row))
	}

	for i, row := range rows {
		for j := range row {
			result[i][j] = row[len(row)-1-j]
		}
	}
	return result
}

func newTile(rows []string) *Tile {
	source := make([][]rune, len(rows))
	for i, row := range rows {
		source[i] = []rune(row)
	}

	rotate90 := rotate(source)
	rotate180 := rotate(rotate90)
	rotate270 := rotate(rotate180)

	rotations := [][][]rune{source, rotate90, rotate180, rotate270}

	variants := make([][][]rune, 0)

	for _, bitmap := range rotations {
		variants = append(variants, bitmap)
		variants = append(variants, flipH(bitmap))
		variants = append(variants, flipV(bitmap))
	}

	return &Tile{variants}
}

func main() {
	file, _ := ioutil.ReadFile("data.txt")
	tiles := make(map[string]*Tile)
	for _, tileText := range strings.Split(string(file), "\n\n") {
		tileRows := strings.Split(tileText, "\n")
		id := strings.TrimPrefix(strings.TrimSuffix(tileRows[0], ":"), "Tile ")
		tiles[id] = newTile(tileRows[1:])
	}

	top := make(map[string][]int)
	bottom := make(map[string][]int)

	for id1, tile1 := range tiles {
		for id2, tile2 := range tiles {
			if id1 != id2 {
				i, j, match := verticalStitch(tile1, tile2)
				if match {
					// if _, ok := matches[id1]; ok {
					// 	matches[id1] = append(matches[id1], topBorder)
					// } else {
					// 	matches[id1] = make([][]rune, 0)
					// }
					// if _, ok := matches[id2]; ok {
					// 	matches[id2] = append(matches[id2], bottomBorder)
					// } else {
					// 	matches[id2] = make([][]rune, 0)
					// }

					if _, ok := top[id1]; ok {
						top[id1] = append(top[id1], i)
					} else {
						top[id1] = make([]int, 0)
					}
					if _, ok := bottom[id2]; ok {
						bottom[id2] = append(bottom[id2], j)
					} else {
						bottom[id2] = make([]int, 0)
					}
					// fmt.Println(tile1.id, tile2.id, i, j)
				}
			}
		}
	}
	fmt.Println("Done")
}
