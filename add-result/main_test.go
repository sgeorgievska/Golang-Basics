package main

import "testing"

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResult = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	//iterate over every obj in AddResult
	for _, test := range addResult {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Received")
		}
	}
}
