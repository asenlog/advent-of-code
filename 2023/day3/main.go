package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(data []string) {
	totalSum := 0

	gameMatrix := make([][]string, len(data))
	for i, row := range data {
		gameMatrix[i] = strings.Split(row, "")
	}

	for _, v := range gameMatrix {
		fmt.Println(v)
	}

	totalSum = sumContiguousNumbers(gameMatrix)

	fmt.Println("Part 1:", totalSum)
}

func sumContiguousNumbers(grid [][]string) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	var totalSum int

	for i, row := range grid {
		for j, symbol := range row {
			if isNumber(symbol) && hasAdjacentSymbol(grid, i, j) && !visited[i][j] {
				contiguousNumber := getContiguousNumber(grid, i, j, visited)
				number, err := strconv.Atoi(contiguousNumber)
				if err == nil {
					fmt.Printf("Number at (%d, %d): %s\n", i+1, j+1, contiguousNumber)
					totalSum += number
				}
			}
		}
	}

	return totalSum
}

func getContiguousNumber(grid [][]string, i, j int, visited [][]bool) string {
	var num string

	// Check left
	for y := j; y >= 0 && isNumber(grid[i][y]) && !visited[i][y]; y-- {
		num = grid[i][y] + num
		visited[i][y] = true
	}

	// Check right
	for y := j + 1; y < len(grid[i]) && isNumber(grid[i][y]) && !visited[i][y]; y++ {
		num = num + grid[i][y]
		visited[i][y] = true
	}

	return num
}

func hasAdjacentSymbol(grid [][]string, i, j int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			newI, newJ := i+x, j+y
			if isValidIndex(grid, newI, newJ) && grid[newI][newJ] != "." && !isNumber(grid[newI][newJ]) {
				return true
			}
		}
	}
	return false
}

func isNumber(symbol string) bool {
	return symbol >= "0" && symbol <= "9"
}

func isValidIndex(grid [][]string, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day3_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}

	part1(data)
}
