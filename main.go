package main

import (
	"fmt"
	"math"
	"math-skills/src/shared"
	"math-skills/src/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	utils.CheckArgs()

	// Reads the file and splits the content of the file into words
	var ints []int
	test := utils.ReadFile(os.Args[1])
	ttt := strings.Fields(test)

	// Iterates, converts the words to int
	for _, v := range ttt {
		vint, err := strconv.Atoi(v)
		utils.CheckErr(err)
		ints = append(ints, vint)
	}

	// Adds each integer v to the sum
	sum := 0
	for _, v := range ints {
		sum += v
	}

	sort.Ints(ints) // sorts the slices
	// calculates the average
	shared.Average = int(math.Round(float64(sum) / float64(len(ints))))
	// calculates the median
	shared.Median = int(math.Round(Median(ints)))
	// prints the results
	fmt.Printf("Average: %d\n", shared.Average)
	fmt.Printf("Median: %v\n", int(math.Round(Median(ints))))
	// calculates the variance, rounds and displays
	fmt.Printf("Variance: %v\n", int(math.Round(Variance(ints, float64(shared.Average)))))
	// calculates standard deviation with Sqrt
	shared.Deviation = int(math.Round(math.Sqrt(float64(shared.Temp))))
	fmt.Printf("Standard Deviation: %v\n", int(math.Round(float64(shared.Deviation))))
}

// calculates the median
func Median(ints []int) float64 {
	sort.Ints(ints)
	length := len(ints)
	if length%2 == 0 {
		mid1 := ints[length/2]
		mid2 := ints[length/2-1]
		return float64(mid1+mid2) / 2.0
	}
	return float64(ints[length/2])
}

// calculates the variance
func Variance(ints []int, average float64) float64 {
	var sum float64

	for _, v := range ints {
		sum += math.Pow(float64(v)-average, 2)
	}
	sum /= float64(len(ints))
	shared.Temp = sum
	return sum
}
