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

	// Lit le fichier et divise le contenu du fichier en mots
	var ints []int
	test := utils.ReadFile(os.Args[1])
	ttt := strings.Fields(test)

	// Parcours, convertit les mots en int
	for _, v := range ttt {
		vint, err := strconv.Atoi(v)
		utils.CheckErr(err)
		ints = append(ints, vint)
	}

	// On ajoute chaque entier v à la somme
	sum := 0
	for _, v := range ints {
		sum += v
	}

	sort.Ints(ints) // trie les slices
	// calcul de la moyenne
	shared.Average = int(math.Round(float64(sum) / float64(len(ints))))
	// calcul médiane
	shared.Median = int(math.Round(Median(ints)))
	// print les résultats
	fmt.Printf("Average: %d\n", shared.Average)
	fmt.Printf("Median: %v\n", int(math.Round(Median(ints))))
	// calcul de la variance, arrondit et affiche
	fmt.Printf("Variance: %v\n", int(math.Round(Variance(ints, float64(shared.Average)))))
	// calcul écart type avec Sqrt
	shared.Deviation = int(math.Round(math.Sqrt(float64(shared.Temp))))
	fmt.Printf("Standard Deviation: %v\n", int(math.Round(float64(shared.Deviation))))
}

// calcul de la médiane
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

// calcul de la variance
func Variance(ints []int, average float64) float64 {
	var sum float64

	for _, v := range ints {
		sum += math.Pow(float64(v)-average, 2)
	}
	sum /= float64(len(ints))
	shared.Temp = sum
	return sum
}
