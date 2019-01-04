package main

import "testing"

func TestSum(t *testing.T) {
	expected := 3
	assert := sum(1, 2)
	if expected != assert {
		t.Fatalf("sum() = %v, want %v", assert, expected)
	}
}
