package main

import "testing"

func TestTemplate(t *testing.T) {
	actual := Template(12, "気温", 22.4)
	expect := "12時の気温は22.4"

	if actual != expect {
		t.Fatalf("failed test. actual = [%#v], expect = [%#v]", actual, expect)
	}
}
