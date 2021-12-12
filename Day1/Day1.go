package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDepths() []int {
	path := "input.txt"
	file, _ := os.ReadFile(path)
	text := strings.Split(strings.TrimRight(string(file), "\n"), "\n")

	depths := make([]int, len(text))
	for i, s := range text {
		depths[i], _ = strconv.Atoi(s)
	}
	return depths
}

func countIncrements(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			count += 1
		}
	}
	return count
}

func countWindowIncrements(depths []int) int {
	count := 0
	for i := 3; i < len(depths); i++ {
		//i-1 and i-2 would be on both sides so can be ignored
		if depths[i-3] < depths[i] {
			count += 1
		}
	}
	return count
}

func main() {
	depths := getDepths()
	incrementCount := countIncrements(depths)
	windowedIncrementCount := countWindowIncrements(depths)

	fmt.Println("The total number of values higher then the last is", incrementCount)
	fmt.Println("The total number of sets of three values higher then the last is", windowedIncrementCount)
}
