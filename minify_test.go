package main

import "testing"
import "reflect"

func TestMinify(t *testing.T) {
	tests := []struct {
		poly   []int
		result []int
	}{
		{[]int{0, 6, 4}, []int{0, 3, 2}},
		{[]int{12, 6, 24}, []int{2, 1, 4}},
	}
	for _, test := range tests {
		result := minify(test.poly)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Testing %v: Expected %v, got %v", test.poly, test.result, result)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{24, 27, 3},
		{0, 72, 72},
		{256, 1024, 256},
	}
	for _, test := range tests {
		result := gcd(test.a, test.b)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Expected %v, got %v", test.result, result)
		}
	}
}
