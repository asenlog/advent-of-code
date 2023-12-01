package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day8_input")
	if err != nil {
		log.Fatal("failed to open file")
	}

	data := strings.Split(string(input), "\n")
	if err != nil {
		log.Fatalf("failed split the input lines")
	}

	part1(data)
	part2(data)
}

func countVisibleTrees(treeMap [][]int32) int {
	// Initialize a count of visible trees to 0
	visibleTrees := 0

	// Iterate through each tree in the grid
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			// Check if the tree is on the edge of the grid
			if i == 0 || j == 0 || i == len(treeMap)-1 || j == len(treeMap[i])-1 {
				// If it is, it is visible and can be counted
				visibleTrees++
				continue
			}

			// Check the trees in the same row and column
			// If any are taller than the current tree, it is not visible and can be skipped
			for k := 0; k < len(treeMap); k++ {
				if treeMap[i][k] > treeMap[i][j] || treeMap[k][j] > treeMap[i][j] {
					continue
				}
			}

			// If the tree is not on the edge and no taller trees are in its row or column,
			// it is visible and can be counted
			visibleTrees++
		}
	}

	return visibleTrees
}

func part1(data []string) {
	treeMap := make([][]int32, len(data))

	for i, line := range data {
		treeMap[i] = make([]int32, len(line))
		for _, char := range line {
			treeMap[i] = append(treeMap[i], char)
		}
	}

	//fmt.Println(treeMap)
	visibleTrees := countVisibleTrees(treeMap)
	fmt.Println(visibleTrees)
}

func part2(data []string) {

}
