package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day5_input")
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

func createTheStack(data []string) map[int]*Stack {
	stack := make(map[int]*Stack, 0)
	for _, line := range data {
		if line == "" || !strings.Contains(line, "[") {
			return stack
		}

		for j := 1; j <= len(line); j += 4 {
			if line[j] == ' ' {
				continue
			}

			stackNum := j/4 + 1
			if _, ok := stack[stackNum]; !ok {
				stack[stackNum] = &Stack{}
			}

			//prepend data to the stack
			*stack[stackNum] = append([]string{string(line[j])}, *stack[stackNum]...)
		}
	}

	return stack
}

// 1st Number: moves 2nd Number: From Stack 3rd Number: To Stack
func parseMoves(input string) []int {
	var result []int
	data := strings.Split(input, " ")
	// fmt.Println(data)

	moves, _ := strconv.Atoi(data[1])
	result = append(result, moves)

	fromStack, _ := strconv.Atoi(data[3])
	result = append(result, fromStack)

	toStack, _ := strconv.Atoi(data[5])
	result = append(result, toStack)

	return result
}

func part1(data []string) {
	stack := createTheStack(data)

	for _, line := range data {
		if !strings.HasPrefix(line, "move") {
			continue
		}

		instructions := parseMoves(line)
		// fmt.Println(instructions)

		for i := 0; i < instructions[0]; i++ {
			s := stack[instructions[1]]
			element, ok := s.Pop()
			if ok {
				*stack[instructions[2]] = append(*stack[instructions[2]], element)
			} else {
				log.Fatalf("failed to pop from stack %d", instructions[1])
			}
		}
	}

	fmt.Print("Part 1: ")
	for i := 1; i <= 9; i++ {
		s := *stack[i]
		fmt.Printf("%s", s[len(s)-1])
	}
	fmt.Println()
}

func part2(data []string) {
	stack := createTheStack(data)

	for _, line := range data {
		if !strings.HasPrefix(line, "move") {
			continue
		}

		instructions := parseMoves(line)
		// fmt.Println(instructions)
		var orderStack []string
		for i := 0; i < instructions[0]; i++ {
			s := stack[instructions[1]]
			element, ok := s.Pop()
			if ok {
				orderStack = append([]string{element}, orderStack...)
			} else {
				log.Fatalf("failed to pop from stack %d", instructions[1])
			}
		}

		*stack[instructions[2]] = append(*stack[instructions[2]], orderStack...)
	}

	fmt.Print("Part 2: ")
	for i := 1; i <= 9; i++ {
		s := *stack[i]
		fmt.Printf("%s", s[len(s)-1])
	}
	fmt.Println()
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
