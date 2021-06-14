package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		got, want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{45, -9, 25, 30, 1}, []int{-9, 1, 25, 30, 45}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2}, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.got, tt.want)
		t.Run(testname, func(t *testing.T) {
			BubbleSort(tt.got)
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

