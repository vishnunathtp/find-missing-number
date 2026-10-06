package main

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}

	for _, c := range cases {
		got := MissingNumber(c.nums)
		if got != c.want {
			t.Errorf("MissingNumber(%v) = %d; want %d", c.nums, got, c.want)
		}
	}
}
