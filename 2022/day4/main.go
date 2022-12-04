package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day4_input")
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

func isSubset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func isPartialSubset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count >= 1 {
			return true
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func part1(data []string) {
	var sum int
	for _, line := range data {
		numbers := strings.Split(line, ",")
		strSet1 := strings.Split(numbers[0], "-")
		strSet2 := strings.Split(numbers[1], "-")

		set1Start, err := strconv.Atoi(strSet1[0])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set1End, err := strconv.Atoi(strSet1[1])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set2Start, err := strconv.Atoi(strSet2[0])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set2End, err := strconv.Atoi(strSet2[1])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		var set1, set2 []int
		for i := set1Start; i <= set1End; i++ {
			set1 = append(set1, i)
		}

		for i := set2Start; i <= set2End; i++ {
			set2 = append(set2, i)
		}

		if isSubset(set1, set2) || isSubset(set2, set1) {
			sum++
		}
	}

	fmt.Println("Pairs fully contain one another: ", sum)
}

func part2(data []string) {
	var sum int
	for _, line := range data {
		numbers := strings.Split(line, ",")
		strSet1 := strings.Split(numbers[0], "-")
		strSet2 := strings.Split(numbers[1], "-")

		set1Start, err := strconv.Atoi(strSet1[0])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set1End, err := strconv.Atoi(strSet1[1])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set2Start, err := strconv.Atoi(strSet2[0])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		set2End, err := strconv.Atoi(strSet2[1])
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		var set1, set2 []int
		for i := set1Start; i <= set1End; i++ {
			set1 = append(set1, i)
		}

		for i := set2Start; i <= set2End; i++ {
			set2 = append(set2, i)
		}

		if isPartialSubset(set1, set2) || isPartialSubset(set2, set1) {
			sum++
		}
	}

	fmt.Println("Pairs partilly contain one another: ", sum)
}
