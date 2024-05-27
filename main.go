package main

import "sort"

func main() {

	//coins := []int{5, 7, 1, 1, 2, 3, 22}
	//coins := []int{1, 1, 1, 1, 1}
	coins := []int{87}

	result := NonConstructibleChange(coins)
	println(result)
}

func NonConstructibleChange(coins []int) int {
	minimumChange := 0

	if len(coins) == 1 {
		if coins[0] > 1 {
			return 1
		}
		if coins[0] == 1 {
			return 2
		}
		return coins[0]
	}
	sort.Ints(coins)
	for i := 0; i < len(coins); i++ {

		if minimumChange == 0 {
			minimumChange = coins[i]
			continue
		}

		if coins[i] > minimumChange+1 {
			return minimumChange + 1
		}
		minimumChange = minimumChange + coins[i]

	}
	return minimumChange + 1
}
