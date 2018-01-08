package main

import "testing"

func TestExtraction(t *testing.T) {
	actual := Extraction("パタトクカシーー")
	expect := "パトカー"

	if expect != actual {
		t.Fatalf("failed test. actual = [%#v], expect = [%#v]", actual, expect)
	}
}
