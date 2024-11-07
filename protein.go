package main

import (
	"errors"
	"fmt"
	//"math"
)

var linkW = [][]int{
	{-1, 1, 1, 0},
	{1, -1, -2, 0},
	{1, -2, 1, 0},
	{0, 0, 0, -1},
}

var unlinkW = [][]int{
	{1, -3, -1, 0},
	{-3, 2, 1, 0},
	{-1, 1, -1, 0},
	{0, 0, 0, -1},
}

var proteinDic = map[rune]int{'A': 0, 'B': 1, 'C': 2, 'T': 3}
var directionDic = map[rune][2]int{'L': {-1, 0}, 'U': {0, 1}, 'R': {1, 0}, 'D': {0, -1}}

func divider(answer string) (string, string) {
	mid := (len(answer) + 1) / 2
	return answer[:mid], answer[mid:]
}

func proteins2Indices(proteins string) []int {
	var indices []int
	for _, protein := range proteins {
		indices = append(indices, proteinDic[protein])
	}
	return indices
}

func direction2Vec(direction rune) [2]int {
	return directionDic[direction]
}

func calculator(answer string, view bool) (int, error) {
	proteins, directions := divider(answer)
	proteinIndices := proteins2Indices(proteins)
	length := len(proteins)

	x, y := 0, 0
	points := [][2]int{{x, y}}
	xmin, xmax, ymin, ymax := 0, 0, 0, 0

	for i := 0; i < length-1; i++ {
		vec := direction2Vec(rune(directions[i]))
		x += vec[0]
		y += vec[1]
		points = append(points, [2]int{x, y})
		xmin, xmax = min(xmin, x), max(xmax, x)
		ymin, ymax = min(ymin, y), max(ymax, y)
	}

	n_x := xmax - xmin + 1
	n_y := ymax - ymin + 1

	map_ := make([][]int, n_y+2)
	for i := range map_ {
		map_[i] = make([]int, n_x+2)
	}

	energy := 0

	for i, point := range points {
		x, y := point[0]-xmin+1, point[1]-ymin+1
		pIndex := proteinIndices[i]
		if map_[y][x] == 0 {
			map_[y][x] = i + 1
		} else {
			return 0, errors.New("error! wrong structure")
		}
		for _, dr := range "LURD" {
			vec := direction2Vec(rune(dr))
			x_, y_ := x+vec[0], y+vec[1]
			if map_[y_][x_] != 0 {
				pNeighbor := proteinIndices[map_[y_][x_]-1]
				if map_[y][x]-map_[y_][x_] == 1 {
					energy += linkW[pIndex][pNeighbor]
				} else {
					energy += unlinkW[pIndex][pNeighbor]
				}
			}
		}
	}

	if view {
		drawProtein(answer)
	}
	return energy, nil
}
func drawProtein(answer string) {
	fmt.Println(answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func drawProtein(answer string) {
	proteins, directions := divider(answer)
	proteinIndices := proteins2Indices(proteins)
	length := len(proteins)
	location := [2]int{0, 0}

	xs := make([]int, length)
	ys := make([]int, length)
	for i, _ := range proteinIndices {
		x, y := location[0], location[1]
		xs[i] = x
		ys[i] = y
		if i != length-1 {
			vec := direction2Vec(rune(directions[i]))
			location[0] += vec[0]
			location[1] += vec[1]
		}
	}

	min_, max_ := math.MaxInt32, math.MinInt32
	for _, val := range append(xs, ys...) {
		if val < min_ {
			min_ = val
		}
		if val > max_ {
			max_ = val
		}
	}
	size := max_ - min_

	// Placeholder for visualization - in Go, you may need to use a library like Gonum plot or export data for plotting in another environment.
	fmt.Println("Plotting functionality would go here.")
}*/
/*
func main() {
	answer := "AAAAAAAAAAAAAAAACBCBCBCBCBCBCBCBCBCBBCBCBCBCBCBCBCBCBCBBCBCBCBCBCBCBCBCBCBCBBCBCBCBCBCBCBCBCBCAAAAAAAAAAAAAAAAAURURURURURURURURDDLDLDLDLDLDLDLDLLUUURURURURURURURURUULDLDLDLDLDLDLDLDLDLLUUURURURURURURURURRDLDLDLDLDLDLDLDLD" // Example input
	for i:=0;i<100000;i++{
	energy, err := calculator(answer, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(i, ">> Energy:", energy)
	}}
}*/
