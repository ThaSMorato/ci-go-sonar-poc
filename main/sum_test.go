package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2,3)

	if result != 5 {
		t.Error("The result should be 5")
	}
}
