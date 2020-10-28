package main

import "fmt"

func main() {
	troco := troco(2.89)
	fmt.Println(troco)
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
