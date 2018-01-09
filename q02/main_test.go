package main

import "testing"

func TestJoinInTurn(t *testing.T) {
	actual := JoinInTurn("パトカー", "タクシー")
	expect := "パタトクカシーー"

	if expect != actual {
		t.Fatalf("failed test. actual = [%#v], expect = [%#v]", actual, expect)
	}
}
