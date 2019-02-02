package main

import (
	"testing"
)

func TestSample1(t *testing.T) {
	file := "testfiles/sample1.txt"

	err := replace(file)
	if err != nil {
		t.Fatal(err)
	}
}
