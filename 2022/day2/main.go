package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CalculateWinner(p1, p2 string) int64 {
	if p1 == "A" {
		if p2 == "X" {
			return 1 + 3
		}

		if p2 == "Y" {
			return 2 + 6
		}

		if p2 == "Z" {
			return 3 + 0
		}
	}

	if p1 == "B" {
		if p2 == "X" {
			return 1 + 0
		}

		if p2 == "Y" {
			return 2 + 3
		}

		if p2 == "Z" {
			return 3 + 6
		}
	}

	if p1 == "C" {
		if p2 == "X" {
			return 1 + 6
		}

		if p2 == "Y" {
			return 2 + 0
		}

		if p2 == "Z" {
			return 3 + 3
		}
	}

	log.Fatalf("Unrecognized entry")
	return 0
}

func CalculateMove(p1, p2 string) string {
	if p1 == "A" {
		if p2 == "X" {
			return "Z" //Z
		}

		if p2 == "Y" {
			return "X" // X
		}

		if p2 == "Z" {
			return "Y" // Y
		}
	}

	if p1 == "B" {
		if p2 == "X" {
			return "X" // X
		}

		if p2 == "Y" {
			return "Y" // Y
		}

		if p2 == "Z" {
			return "Z" //Z
		}
	}

	if p1 == "C" {
		if p2 == "X" {
			return "Y" // Y
		}

		if p2 == "Y" {
			return "Z" //Z
		}

		if p2 == "Z" {
			return "X" // X
		}
	}

	log.Fatalf("Unrecognized entry")
	return ""
}

func part1(data []string) {
	var sum int64
	for _, line := range data {
		roundsData := strings.Split(line, " ")
		//fmt.Printf("line is: %s \n", roundsData[1])
		player1 := strings.TrimSpace(roundsData[0])
		player2 := strings.TrimSpace(roundsData[1])

		sum += CalculateWinner(player1, player2)

	}

	fmt.Printf("Final Score is: %d", sum)

}

func part2(data []string) {
	var sum int64
	for _, line := range data {
		roundsData := strings.Split(line, " ")
		//fmt.Printf("line is: %s \n", roundsData[1])
		player1 := strings.TrimSpace(roundsData[0])
		player2 := strings.TrimSpace(roundsData[1])

		move := CalculateMove(player1, player2)
		sum += CalculateWinner(player1, move)

	}

	fmt.Printf("Final Score is: %d \n", sum)
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day2_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}

	part1(data)
	fmt.Println()
	part2(data)
}
