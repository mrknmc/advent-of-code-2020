package main

import (
	"bufio"
	"fmt"
	"os"
)

type State int

const (
	Inactive State = 0
	Active         = 1
)

func hyperNeighborLocations() [][]int {
	locations := make([][]int, 0)
	values := []int{-1, 0, 1}
	for _, x := range values {
		for _, y := range values {
			for _, z := range values {
				for _, w := range values {
					locations = append(locations, []int{x, y, z, w})
				}
			}
		}
	}
	return locations
}

func neighborLocations() [][]int {
	locations := make([][]int, 0)
	values := []int{-1, 0, 1}
	for _, x := range values {
		for _, y := range values {
			for _, z := range values {
				locations = append(locations, []int{x, y, z, 0})
			}
		}
	}
	return locations
}

type HyperCube struct {
	x, y, z, w int
	states     *[]State
}

func (cube *HyperCube) isActive() bool {
	return (*cube.states)[len(*cube.states)-1] == Active
}

type HyperCubeMap struct {
	data  *map[string]*HyperCube
	hyper bool
}

func (cubes *HyperCubeMap) countActiveNeighbors(cube *HyperCube, cycle int) int {
	activeCount := 0
	for _, neighbor := range cubes.neighbors(cube, cycle) {
		if (*neighbor.states)[cycle] == Active {
			activeCount++
		}
	}
	return activeCount
}

func (cubes *HyperCubeMap) neighbors(cube *HyperCube, cycle int) []*HyperCube {
	neighbors := make([]*HyperCube, 0)
	var locations [][]int
	if cubes.hyper {
		locations = hyperNeighborLocations()
	} else {
		locations = neighborLocations()
	}
	for _, loc := range locations {
		x, y, z, w := loc[0], loc[1], loc[2], loc[3]
		// assume it's pre-computed; skip self
		if x != 0 || y != 0 || z != 0 || w != 0 {
			neighbor, _ := cubes.get(cube.x+x, cube.y+y, cube.z+z, cube.w+w)
			neighbors = append(neighbors, neighbor)
		}

	}
	return neighbors
}

func (cubes *HyperCubeMap) cycle(cube *HyperCube, cycle int) {
	activeCount := cubes.countActiveNeighbors(cube, cycle-1)
	lastState := (*cube.states)[cycle-1]

	if lastState == Active && (activeCount == 2 || activeCount == 3) {
		*cube.states = append(*cube.states, Active)
	} else if lastState == Active {
		*cube.states = append(*cube.states, Inactive)
	} else if lastState == Inactive && activeCount == 3 {
		*cube.states = append(*cube.states, Active)
	} else if lastState == Inactive {
		*cube.states = append(*cube.states, Inactive)
	} else {
		panic("Impossible state")
	}
}

func (cubes *HyperCubeMap) computeNeighbors(cycle int) {
	for _, cube := range cubes.getAll() {
		var locations [][]int
		if cubes.hyper {
			locations = hyperNeighborLocations()
		} else {
			locations = neighborLocations()
		}
		for _, loc := range locations {
			x, y, z, w := loc[0], loc[1], loc[2], loc[3]
			if _, ok := cubes.get(cube.x+x, cube.y+y, cube.z+z, cube.w+w); !ok {
				cubes.add(newHyperCube(cube.x+x, cube.y+y, cube.z+z, cube.w+w, Inactive, cycle))
			}
		}
	}
}

func (cubes *HyperCubeMap) getAll() []*HyperCube {
	result := make([]*HyperCube, 0)
	for _, cube := range *cubes.data {
		result = append(result, cube)
	}
	return result
}

func (cubes *HyperCubeMap) get(x int, y int, z int, w int) (*HyperCube, bool) {
	loc := fmt.Sprintf("%d.%d.%d.%d", x, y, z, w)
	cube, ok := (*cubes.data)[loc]
	return cube, ok
}

func (cubes *HyperCubeMap) add(cube *HyperCube) {
	loc := fmt.Sprintf("%d.%d.%d.%d", cube.x, cube.y, cube.z, cube.w)
	(*cubes.data)[loc] = cube
}

func (cubes *HyperCubeMap) delete(cube *HyperCube) {
	loc := fmt.Sprintf("%d.%d.%d.%d", cube.x, cube.y, cube.z, cube.w)
	delete(*cubes.data, loc)
}

func newHyperCube(x int, y int, z int, w int, state State, cycle int) *HyperCube {
	// fill previous states with inactive
	states := make([]State, cycle+1)
	for i := 0; i < cycle; i++ {
		states[i] = Inactive
	}
	states[cycle] = state

	return &HyperCube{x, y, z, w, &states}
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for fscanner.Scan() {
		lines = append(lines, fscanner.Text())
	}

	data := make(map[string]*HyperCube)
	cubes := &HyperCubeMap{&data, false}

	for x, line := range lines {
		for y, char := range line {
			if char == '#' {
				cubes.add(newHyperCube(x, y, 0, 0, Active, 0))
			} else {
				cubes.add(newHyperCube(x, y, 0, 0, Inactive, 0))
			}
		}
	}

	for i := 1; i < 7; i++ {
		// new neighbors were inactive in last cycle
		cubes.computeNeighbors(i - 1)
		candidates := cubes.getAll()
		// new neighbors of neighbors were inactive in last cycle
		// and are inactive in current cycle
		cubes.computeNeighbors(i)
		for _, cube := range candidates {
			cubes.cycle(cube, i)
		}
	}
	countPart1 := 0
	for _, cube := range cubes.getAll() {
		if cube.isActive() {
			countPart1++
		}
	}
	fmt.Println("Part 1 ", countPart1)

	hyperdata := make(map[string]*HyperCube)
	hypercubes := &HyperCubeMap{&hyperdata, true}
	for x, line := range lines {
		for y, char := range line {
			if char == '#' {
				hypercubes.add(newHyperCube(x, y, 0, 0, Active, 0))
			} else {
				hypercubes.add(newHyperCube(x, y, 0, 0, Inactive, 0))
			}
		}
	}

	for i := 1; i < 7; i++ {
		// new neighbors were inactive in last cycle
		hypercubes.computeNeighbors(i - 1)
		candidates := hypercubes.getAll()
		// new neighbors of neighbors were inactive in last cycle
		// and are inactive in current cycle
		hypercubes.computeNeighbors(i)
		for _, cube := range candidates {
			hypercubes.cycle(cube, i)
		}
	}
	countPart2 := 0
	for _, cube := range hypercubes.getAll() {
		if cube.isActive() {
			countPart2++
		}
	}

	fmt.Println("Part 2 ", countPart2)
}
