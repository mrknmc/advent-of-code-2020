package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	Forward = 'F'
	Right   = 'R'
	Left    = 'L'
	West    = 'W'
	East    = 'E'
	South   = 'S'
	North   = 'N'
)

type Action struct {
	name rune
	val  float64
}

type Waypoint struct {
	location []float64
}

type Ferry struct {
	location []float64
	angle    float64
}

func nameToVectorMap() map[rune][]float64 {
	result := make(map[rune][]float64)
	result['N'] = []float64{0, 1}
	result['S'] = []float64{0, -1}
	result['E'] = []float64{1, 0}
	result['W'] = []float64{-1, 0}
	return result
}

func getRotationMatrix(angle float64) [][]float64 {
	return [][]float64{
		[]float64{math.Round(math.Cos(angle / 180 * math.Pi)), -math.Round(math.Sin(angle / 180 * math.Pi))},
		[]float64{math.Round(math.Sin(angle / 180 * math.Pi)), math.Round(math.Cos(angle / 180 * math.Pi))},
	}
}

func getUnitVector(angle float64) []float64 {
	return []float64{math.Round(math.Cos(angle / 180 * math.Pi)), math.Round(math.Sin(angle / 180 * math.Pi))}
}

func vectorAdd(vector1 []float64, vector2 []float64) []float64 {
	return []float64{vector1[0] + vector2[0], vector1[1] + vector2[1]}
}

func vectorDiff(vector1 []float64, vector2 []float64) []float64 {
	return []float64{vector1[0] - vector2[0], vector1[1] - vector2[1]}
}

func vectorMult(vector1 []float64, mult float64) []float64 {
	return []float64{vector1[0] * mult, vector1[1] * mult}
}

func matrixMult(vector []float64, matrix [][]float64) []float64 {
	return []float64{
		matrix[0][0]*vector[0] + matrix[0][1]*vector[1],
		matrix[1][0]*vector[0] + matrix[1][1]*vector[1],
	}
}

func parseFile(file *os.File) []Action {
	fscanner := bufio.NewScanner(file)
	actions := make([]Action, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		val, _ := strconv.ParseFloat(line[1:], 64)
		action := Action{rune(line[0]), val}
		actions = append(actions, action)
	}
	return actions
}

func main() {
	file, _ := os.Open("data.txt")
	actions := parseFile(file)
	nameToVector := nameToVectorMap()
	ferry := Ferry{[]float64{0.0, 0.0}, 0.0}
	for _, action := range actions {
		if vec, ok := nameToVector[action.name]; ok {
			ferry.location = vectorAdd(ferry.location, vectorMult(vec, action.val))
		} else if action.name == 'R' || action.name == 'L' {
			angle := action.val
			if action.name == 'R' {
				angle = 360 - angle
			}
			ferry.angle = math.Mod(ferry.angle+angle, 360)
		} else if action.name == 'F' {
			ferry.location = vectorAdd(ferry.location, vectorMult(getUnitVector(ferry.angle), action.val))
		} else {
			panic("Impossible state")
		}
	}
	fmt.Println(math.Abs(ferry.location[0]) + math.Abs(ferry.location[1]))

	// reset ferry
	ferry.location = []float64{0.0, 0.0}
	waypoint := Waypoint{[]float64{10.0, 1.0}}

	for _, action := range actions {
		if vec, ok := nameToVector[action.name]; ok {
			waypoint.location = vectorAdd(waypoint.location, vectorMult(vec, action.val))
		} else if action.name == 'R' || action.name == 'L' {
			angle := action.val
			if action.name == 'R' {
				angle = 360 - angle
			}
			newDiff := matrixMult(vectorDiff(waypoint.location, ferry.location), getRotationMatrix(angle))
			waypoint.location = vectorAdd(ferry.location, newDiff)
		} else if action.name == 'F' {
			diff := vectorDiff(waypoint.location, ferry.location)
			ferry.location = vectorAdd(ferry.location, vectorMult(diff, action.val))
			waypoint.location = vectorAdd(ferry.location, diff)
		} else {
			panic("Impossible state")
		}
	}
	fmt.Println(math.Abs(ferry.location[0]) + math.Abs(ferry.location[1]))
}
