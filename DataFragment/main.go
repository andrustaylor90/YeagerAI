package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// simpleHash generates a fixed-length hash (30 characters) from the input string
func simpleHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:30]
}

// reconstructData reassembles the original data from its fragments
func reconstructData(fragments map[int]map[string]string) string {
	// Create a slice of keys to sort the fragments
	keys := make([]int, 0, len(fragments))
	for key := range fragments {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Reconstruct the data and verify hashes
	var reconstructedData strings.Builder
	for _, key := range keys {
		fragment := fragments[key]
		data := fragment["data"]
		hash := fragment["hash"]

		// Verify the integrity of the fragment
		if simpleHash(data) != hash {
			return "Error: Data integrity verification failed."
		}

		reconstructedData.WriteString(data)
	}

	return reconstructedData.String()
}

func main() {
	// Example fragments
	fragments := map[int]map[string]string{
		1: {"data": "Hello", "hash": simpleHash("Hello")},
		2: {"data": "World", "hash": simpleHash("World")},
		3: {"data": "!", "hash": simpleHash("!")},
	}

	originalData := reconstructData(fragments)
	fmt.Println("Fragments:", fragments)
	fmt.Println("Original Data:", originalData)
}
