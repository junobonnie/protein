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
	size := len(answer) + 2
	map_ := make([][]int, size)
	for i := range map_ {
		map_[i] = make([]int, size)
		for j := range map_[i] {
			map_[i][j] = -1
		}
	}
	location := [2]int{length, length}
	energy := 0
	pLast := -1

	for i, pIndex := range proteinIndices {
		x, y := location[0], location[1]
		if map_[y][x] == -1 {
			map_[y][x] = pIndex
		} else {
			return 0, errors.New("error! wrong structure!")
		}
		if i != 0 {
			energy += linkW[pIndex][pLast] - unlinkW[pIndex][pLast]
			for _, dr := range "LURD" {
				vec := direction2Vec(rune(dr))
				x_, y_ := location[0]+vec[0], location[1]+vec[1]
				pNeighbor := map_[y_][x_]
				if pNeighbor != -1 {
					energy += unlinkW[pIndex][pNeighbor]
				}
			}
		}
		if i != length-1 {
			vec := direction2Vec(rune(directions[i]))
			location[0] += vec[0]
			location[1] += vec[1]
		}
		pLast = pIndex
	}
	if view {
		drawProtein(answer)
	}
	return energy, nil
}
func drawProtein(answer string) {
	fmt.Println(answer)
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
