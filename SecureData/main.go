package main

import (
	"fmt"
	"math"
	"sort"
)

// Function to check if a given maximum risk is feasible
func isFeasible(dataCenters []int, fragments int, maxRisk int) bool {
	count := 0
	for _, risk := range dataCenters {
		f := 0
		for ; f <= fragments; f++ {
			if int(math.Pow(float64(risk), float64(f))) > maxRisk {
				break
			}
		}
		count += f - 1
		if count >= fragments {
			return true
		}
	}
	return false
}

// Function to find the minimized maximum risk
func distributeFragments(dataCenters []int, fragments int) int {
	sort.Ints(dataCenters)
	left := 1
	right := int(math.Pow(float64(dataCenters[len(dataCenters)-1]), float64(fragments)))

	for left < right {
		mid := (left + right) / 2
		if isFeasible(dataCenters, fragments, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	dataCenters := []int{10, 20, 30}
	fragments := 5
	minRisk := distributeFragments(dataCenters, fragments)
	fmt.Printf("Minimized maximum risk: %d\n", minRisk)
}
