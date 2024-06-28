package main

import (
	"testing"
)

func TestFindMinimumLatencyPath(t *testing.T) {
	tests := []struct {
		name             string
		graph            map[string][][2]interface{}
		compressionNodes []string
		source           string
		destination      string
		expected         int
	}{
		{
			name: "basic case",
			graph: map[string][][2]interface{}{
				"A": {{"B", 10}, {"C", 20}},
				"B": {{"D", 15}},
				"C": {{"D", 30}},
				"D": {},
			},
			compressionNodes: []string{"B", "C"},
			source:           "A",
			destination:      "D",
			expected:         17,
		},
		{
			name: "no compression node usage",
			graph: map[string][][2]interface{}{
				"A": {{"B", 10}, {"C", 50}},
				"B": {{"D", 10}},
				"C": {{"D", 10}},
				"D": {},
			},
			compressionNodes: []string{},
			source:           "A",
			destination:      "D",
			expected:         20,
		},
		{
			name: "compression node reduces latency",
			graph: map[string][][2]interface{}{
				"A": {{"B", 10}, {"C", 50}},
				"B": {{"D", 10}},
				"C": {{"D", 10}},
				"D": {},
			},
			compressionNodes: []string{"B"},
			source:           "A",
			destination:      "D",
			expected:         15,
		},
		{
			name: "no path to destination",
			graph: map[string][][2]interface{}{
				"A": {{"B", 10}},
				"B": {{"C", 20}},
				"C": {},
			},
			compressionNodes: []string{"B"},
			source:           "A",
			destination:      "D",
			expected:         -1,
		},
		{
			name: "larger graph with compression",
			graph: map[string][][2]interface{}{
				"A": {{"B", 5}, {"C", 10}},
				"B": {{"D", 10}, {"E", 5}},
				"C": {{"D", 5}},
				"D": {{"F", 5}},
				"E": {{"F", 10}},
				"F": {{"G", 5}},
				"G": {},
			},
			compressionNodes: []string{"B", "C", "D"},
			source:           "A",
			destination:      "G",
			expected:         20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := find_minimum_latency_path(tt.graph, tt.compressionNodes, tt.source, tt.destination)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
