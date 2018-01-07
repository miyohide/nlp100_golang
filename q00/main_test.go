package main

import "testing"

func TestReverse(t *testing.T) {
	actual := Reverse("stressed")
	expect := "desserts"

	if expect != actual {
		t.Fatalf("failed test. actual = [%#v], expect = [%#v]", actual, expect)
	}
}
