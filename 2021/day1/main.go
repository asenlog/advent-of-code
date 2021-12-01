package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type MPoints struct {
	A int64
	B int64
	C int64
}

func (p MPoints) Sum() int64  {
	return p.A + p.B + p.C
}

func Create3WindowPairs(data []string) []MPoints {
	var measurements []MPoints
	currentMeasurement := MPoints{}

	for i := 0; i+3 <= len(data); i++ {
		currentMeasurement.A, _ = strconv.ParseInt(data[i], 10, 64)
		currentMeasurement.B, _ = strconv.ParseInt(data[i+1], 10, 64)
		currentMeasurement.C, _ = strconv.ParseInt(data[i+2], 10, 64)

		measurements = append(measurements, currentMeasurement)
	}

	return measurements
}

func part1(data []string) {
	counter := int64(0)
	currentMeasurement := int64(0)
	previousMeasurement := int64(0)

	for i, measurement := range data {
		currentMeasurement, _  = strconv.ParseInt(measurement,10,64)

		if i == 0 {
			previousMeasurement = currentMeasurement
			continue
		}

		if currentMeasurement > previousMeasurement {
			counter+=1
		}

		previousMeasurement = currentMeasurement
	}

	fmt.Println("Depth increased: ",  counter, " times")
}

func part2(data []string) {
	counter := int64(0)
	previousMeasurement := MPoints{}

	measurements := Create3WindowPairs(data)
	for i, measurement := range measurements {
		if i == 0 {
			previousMeasurement = measurement
			continue
		}

		if measurement.Sum() > previousMeasurement.Sum() {
			counter+=1
		}

		previousMeasurement = measurement
	}

	fmt.Println("Depth increased: ",  counter, " times")
}

func parseInput(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	data, err := parseInput("day1_input")
	if err != nil {
		log.Fatalf("failed to open the file")
	}
	part1(data)
	part2(data)
}
