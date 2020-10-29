package main

import (
	"fmt"
	"math"
)

func main() {
	troco := troco(2.89)
	fmt.Println(troco)

	init := []int{4, 6, 13, 4, 2, 6, 7, 9, 1, 3, 9}
	final := []int{8, 7, 14, 5, 4, 9, 10, 11, 6, 13, 12}
	sdm := sdm(init, final, len(init))
	fmt.Println(sdm)

}

var availableCoins = [5]float64{1., .25, .10, .05, .01}

func troco(value float64) (coins []float64) {
	sum := 0.
	index := 0
	var highestAvailableCoin float64

	for sum != value {
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
	f[0] = math.MinInt8
	i := 0

	for k := 1; k < n; k++ {
		if s[k] > f[i] {
			x = append(x, k)
			i = k
		}
	}

	return
}
