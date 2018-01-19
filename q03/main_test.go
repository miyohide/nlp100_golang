package main

import "testing"

func TestCountWords(t *testing.T) {
	_, actual_counts := CountWords("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	expect_counts := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}

	for i, v := range actual_counts {
		if expect_counts[i] != v {
			t.Fatalf("failed test. actual = [%#v], expect = [%#v]", v, expect_counts[i])
		}
	}
}
