package main

import (
	"fmt"
	"testing"
)

func TestFindian(t *testing.T) {
	var tests = []struct {
		str, want string
	}{
		{"iannn", "Found!"},
		{"Ian", "Found!"},
		{"iuiygaygn", "Found!"},
		{"I d skd a efju N", "Found!"},
		{"ihhhhhn", "Not Found!"},
		{"ina", "Not Found!"},
		{"ina", "Not Found!"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.str, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := findian(tt.str)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

