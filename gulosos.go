package main

import (
	"fmt"
	"math"
)

var count int

func main() {
	count = 0
	troco := troco(2.89)
	fmt.Println("Resultado Troco: ", troco, " | Iterações: ", count)
	fmt.Println()

	count = 0
	init := []int{4, 6, 13, 4, 2, 6, 7, 9, 1, 3, 9}
	final := []int{8, 7, 14, 5, 4, 9, 10, 11, 6, 13, 12}
	sdmRes := sdm(init, final, len(init))
	fmt.Println("Resultado não ordenado: ", sdmRes, " | Iterações: ", count)
	fmt.Println()

	count = 0
	// Com valores ordenados
	initSorted := []int{2, 4, 1, 6, 4, 6, 7, 9, 9, 3, 13}
	finalSorted := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	sdmRes = sdm(initSorted, finalSorted, len(initSorted))
	fmt.Println("Resultado ordenado: ", sdmRes, " | Iterações: ", count)
}

var availableCoins = [5]float64{1., .25, .10, .05, .01}

func troco(value float64) (coins []float64) {
	sum := 0.
	index := 0
	var highestAvailableCoin float64

	for sum != value {
		count++
		highestAvailableCoin = availableCoins[index]

		if highestAvailableCoin+sum > value {
			index = index + 1
			if len(availableCoins) == index {
				return
			}

		} else {
			coins = append(coins, highestAvailableCoin)
			sum = sum + highestAvailableCoin
		}
	}

	return
}

func sdm(s []int, f []int, n int) (x []int) {
	f = append([]int{math.MinInt8}, f...) // adiciona valor negativo ao inicio do vetor
	s = append([]int{math.MinInt8}, s...) // adiciona valor negativo ao inicio do vetor
	i := 0

	for k := 1; k <= n; k++ {
		count++

		if s[k] > f[i] {
			x = append(x, k)
			i = k
		}
	}

	return
}
