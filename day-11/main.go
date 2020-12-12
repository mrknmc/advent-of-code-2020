package main

import (
	"bufio"
	"fmt"
	"os"
)

type State int

const (
	Floor    State = '.'
	Empty    State = 'L'
	Occupied State = '#'
)

func parseFile(file *os.File) [][]State {
	fscanner := bufio.NewScanner(file)
	rows := make([][]State, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		row := make([]State, 0)
		for _, c := range line {
			if c == 'L' {
				row = append(row, Empty)
			} else if c == '.' {
				row = append(row, Floor)
			} else if c == '#' {
				row = append(row, Occupied)
			} else {
				panic("Unknown character")
			}
		}
		rows = append(rows, row)
	}
	return rows
}

func match(rows1 [][]State, rows2 [][]State) bool {
	for i := range rows1 {
		for j := range rows1[i] {
			if rows1[i][j] != rows2[i][j] {
				return false
			}

		}
	}
	return true
}

func copy(rows [][]State) [][]State {
	rowsCopy := make([][]State, len(rows))
	for i := range rows {
		colsCopy := make([]State, len(rows[0]))
		for j := range rows[0] {
			colsCopy[j] = rows[i][j]
		}
		rowsCopy[i] = colsCopy
	}
	return rowsCopy
}

func countVisiblyOccupied(seats [][]State, originX int, originY int, vectorX int, vectorY int) int {
	x := originX + vectorX
	y := originY + vectorY
	for {
		if x == len(seats) || x == -1 || y == len(seats[0]) || y == -1 {
			// out of bounds
			break
		}
		if seats[x][y] == Occupied {
			return 1
		} else if seats[x][y] == Empty {
			return 0
		}

		x += vectorX
		y += vectorY
	}
	// floors all the way
	return 0
}

func countOccupied(seats [][]State, originX int, originY int, vectorX int, vectorY int) int {
	x := originX + vectorX
	y := originY + vectorY
	if x == len(seats) || x == -1 || y == len(seats[0]) || y == -1 {
		// out of bounds
		return 0
	}

	if seats[x][y] == Occupied {
		return 1
	}
	return 0
}

func print(rows [][]State) {
	for _, row := range rows {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println("")
	}
}

func main() {
	file, _ := os.Open("data.txt")
	rows := parseFile(file)
	original := copy(rows)
	colCount := len(rows)
	rowCount := len(rows[0])
	rowsCopy := copy(rows)
	step := 0
	for {
		for x := 0; x < rowCount; x++ {
			for y := 0; y < colCount; y++ {
				vectors := [][]int{
					[]int{0, 1},
					[]int{0, -1},
					[]int{1, 0},
					[]int{1, 1},
					[]int{1, -1},
					[]int{-1, 0},
					[]int{-1, 1},
					[]int{-1, -1},
				}
				occupied := 0
				for _, vec := range vectors {
					occupied += countOccupied(rows, x, y, vec[0], vec[1])
				}
				if rows[x][y] == Empty && occupied == 0 {
					rowsCopy[x][y] = Occupied
				} else if rows[x][y] == Occupied && occupied >= 4 {
					rowsCopy[x][y] = Empty
				} else {
					rowsCopy[x][y] = rows[x][y]
				}
			}
		}
		step++
		temp := rows
		rows = rowsCopy
		rowsCopy = temp
		if match(rows, rowsCopy) {
			break
		}
	}

	count := 0
	for _, row := range rows {
		for _, col := range row {
			if col == Occupied {
				count++
			}
		}
	}
	fmt.Println(count)

	// reset seats
	rows = original

	for {
		for x := 0; x < rowCount; x++ {
			for y := 0; y < colCount; y++ {
				vectors := [][]int{
					[]int{0, 1},
					[]int{0, -1},
					[]int{1, 0},
					[]int{1, 1},
					[]int{1, -1},
					[]int{-1, 0},
					[]int{-1, 1},
					[]int{-1, -1},
				}
				occupied := 0
				for _, vec := range vectors {
					occupied += countVisiblyOccupied(rows, x, y, vec[0], vec[1])
				}
				if rows[x][y] == Empty && occupied == 0 {
					rowsCopy[x][y] = Occupied
				} else if rows[x][y] == Occupied && occupied >= 5 {
					rowsCopy[x][y] = Empty
				} else {
					rowsCopy[x][y] = rows[x][y]
				}
			}
		}
		step++
		temp := rows
		rows = rowsCopy
		rowsCopy = temp
		if match(rows, rowsCopy) {
			break
		}
	}

	count = 0
	for _, row := range rows {
		for _, col := range row {
			if col == Occupied {
				count++
			}
		}
	}
	fmt.Println(count)
}
