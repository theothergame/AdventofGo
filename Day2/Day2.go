package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDirectionsAndDistances() ([]string, []int) {
	path := "input.txt"
	file, _ := os.ReadFile(path)
	text := strings.Split(strings.TrimRight(string(file), "\n"), "\n")

	directions := make([]string, len(text))
	distances := make([]int, len(text))
	for i, s := range text {
		splitValue := strings.Split(s, " ")
		directions[i] = splitValue[0]
		distances[i], _ = strconv.Atoi(splitValue[1])
	}
	return directions, distances
}

func calcSimpleDistance(directions []string, distances []int) int {
	depth := 0
	horizontalPosition := 0
	for i, s := range directions {
		switch {
		case s == "forward":
			horizontalPosition += distances[i]
		case s == "down":
			depth += distances[i]
		case s == "up":
			depth -= distances[i]
		}
	}
	return depth * horizontalPosition
}

func calcAimedDistance(directions []string, distances []int) int {
	depth := 0
	horizontalPosition := 0
	aim := 0
	for i, direction := range directions {
		switch {
		case direction == "forward":
			horizontalPosition += distances[i]
			depth += distances[i] * aim
		case direction == "down":
			aim += distances[i]
		case direction == "up":
			aim -= distances[i]
		}
	}
	return depth * horizontalPosition
}

func main() {
	directions, distances := getDirectionsAndDistances()
	simpleDistance := calcSimpleDistance(directions, distances)
	aimedDistance := calcAimedDistance(directions, distances)

	fmt.Println("The simple manhatten distance traveled is", simpleDistance)
	fmt.Println("The aimed manhatten distance traveled is", aimedDistance)
}
