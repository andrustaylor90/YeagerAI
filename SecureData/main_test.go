package main

import (
	"fmt"
	"testing"
)

func TestDistributeFragments(t *testing.T) {
	tests := []struct {
		dataCenters []int
		fragments   int
		expected    int
	}{
		{[]int{10, 20, 30}, 5, 400},      // Given example
		{[]int{1, 2, 3}, 6, 1},           // All risks are low
		{[]int{10, 20, 30}, 1, 10},       // Only one fragment
		{[]int{2, 4, 8}, 3, 4},           // Fragment count equals data centers
		{[]int{5, 5, 5, 5}, 10, 125},     // Same risk factor
		{[]int{1, 1, 1}, 10, 1},          // Minimal risk factors
		{[]int{100, 200, 300}, 6, 90000}, // High risk factors
		{[]int{1, 10, 100}, 6, 1},        // Mixed risk factors
		{[]int{10, 15, 20}, 8, 3375},     // Various risk factors
		{[]int{1, 2, 3, 4, 5}, 15, 1},    // More fragments than data centers
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("dataCenters: %v, fragments: %d", test.dataCenters, test.fragments), func(t *testing.T) {
			result := distributeFragments(test.dataCenters, test.fragments)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
