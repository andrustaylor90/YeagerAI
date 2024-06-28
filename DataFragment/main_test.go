package main

import "testing"

// Test functions
func TestSimpleHash(t *testing.T) {
	hash1 := simpleHash("Hello")
	hash2 := simpleHash("World")
	hash3 := simpleHash("Hello")

	if len(hash1) != 30 {
		t.Errorf("Expected hash length of 30, but got %d", len(hash1))
	}

	if hash1 == hash2 {
		t.Error("Expected different hashes for different inputs, but got the same")
	}

	if hash1 != hash3 {
		t.Error("Expected the same hashes for the same input, but got different")
	}
}

func TestReconstructData(t *testing.T) {
	fragments := map[int]map[string]string{
		1: {"data": "Hello", "hash": simpleHash("Hello")},
		2: {"data": "World", "hash": simpleHash("World")},
		3: {"data": "!", "hash": simpleHash("!")},
	}
	expected := "HelloWorld!"
	result := reconstructData(fragments)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestReconstructDataWithInvalidHash(t *testing.T) {
	fragments := map[int]map[string]string{
		1: {"data": "Hello", "hash": simpleHash("Hello")},
		2: {"data": "World", "hash": "invalidhash"},
		3: {"data": "!", "hash": simpleHash("!")},
	}
	expected := "Error: Data integrity verification failed."
	result := reconstructData(fragments)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestReconstructDataWithOutOfOrderFragments(t *testing.T) {
	fragments := map[int]map[string]string{
		2: {"data": "World", "hash": simpleHash("World")},
		1: {"data": "Hello", "hash": simpleHash("Hello")},
		3: {"data": "!", "hash": simpleHash("!")},
	}
	expected := "HelloWorld!"
	result := reconstructData(fragments)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
