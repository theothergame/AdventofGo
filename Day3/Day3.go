package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getValues() []string {
	path := "input.txt"
	file, _ := os.ReadFile(path)
	text := strings.Split(strings.TrimRight(string(file), "\n"), "\n")
	return text
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func findMajorityBits(values []string) []int {
	digits := len(values[0])
	majorityValues := make([]int, digits)
	for i := 0; i < digits; i++ {
		sum := 0
		for _, s := range values {
			if string(s[i]) == "1" {
				sum++
			}
		}
		if sum >= len(values)/2 {
			majorityValues[i] = 1
		}
	}
	return majorityValues
}

func binaryIntSliceToInt(binaryIntSlice []int) int {
	result := 0
	digits := len(binaryIntSlice)
	for i := 0; i < digits; i++ {
		if binaryIntSlice[i] == 1 {
			result += powInt(2, digits-i-1)
		}
	}
	return result
}

func calcGammaAndEpsilon(values []string) int {
	majorityBits := findMajorityBits(values)
	digits := len(majorityBits)
	gamma := binaryIntSliceToInt(majorityBits)
	epsilon := powInt(2, digits) - gamma - 1
	return gamma * epsilon
}

func calcO2AndCO2(values []string) int64 {
	digits := len(values[0])
	remainingO2Values := values
	remainingCO2Values := values

	for i := 0; i < digits; i++ {
		var new02Values []string
		majorityBits := findMajorityBits(remainingO2Values)
		for _, remainingO2Value := range remainingO2Values {
			bitValue, _ := strconv.Atoi(string(remainingO2Value[i]))
			if majorityBits[i] == bitValue {
				new02Values = append(new02Values, remainingO2Value)
			}
		}
		remainingO2Values = new02Values
		if len(remainingO2Values) == 1 {
			break
		}
	}

	for i := 0; i < digits; i++ {
		var newCO2Values []string
		majorityBits := findMajorityBits(remainingCO2Values)
		for _, remainingCO2Value := range remainingCO2Values {
			bitValue, _ := strconv.Atoi(string(remainingCO2Value[i]))
			if majorityBits[i] != bitValue {
				newCO2Values = append(newCO2Values, remainingCO2Value)
			}
		}
		remainingCO2Values = newCO2Values
		if len(remainingCO2Values) == 1 {
			break
		}
	}

	cO2Value, _ := strconv.ParseInt(remainingCO2Values[0], 2, 64)
	O2Value, _ := strconv.ParseInt(remainingO2Values[0], 2, 64)
	return O2Value * cO2Value
}

func main() {
	values := getValues()
	gammaAndEpsilon := calcGammaAndEpsilon(values)
	o2AndCo2 := calcO2AndCO2(values)
	fmt.Println("The product of the gamma and epsilon values is", gammaAndEpsilon)
	fmt.Println("The product of the o2 and co2 values is", o2AndCo2)
}
