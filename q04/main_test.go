package main

import "testing"

func TestWordIndex(t *testing.T) {
	actual := WordIndex("Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.")
	expect := map[string]int{"H": 1, "He": 2, "Li": 3, "Be": 4, "B": 5, "C": 6, "N": 7, "O": 8, "F": 9, "Ne": 10, "Na": 11, "Mi": 12, "Al": 13, "Si": 14, "P": 15, "S": 16, "Cl": 17, "Ar": 18, "K": 19, "Ca": 20}

	for k, v := range expect {
		if actual[k] != v {
			t.Fatalf("failed test. actual = [%#v], expect = [%#v]", v, actual[k])
		}
	}
}
