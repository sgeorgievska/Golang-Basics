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

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("test.data")
	if err != nil {
		t.Fatal("Could not open the file")
	}
	if string(data) != "Hey Everyone" {
		t.Fatal("Content is not matching")
	}
}
