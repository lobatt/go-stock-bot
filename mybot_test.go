package main

import (
	"testing"
)

func TestGetSymbol(t *testing.T) {
	cases := []struct{
		text string 
		symbol []string
	} {
		{"check $GOOG", []string{"GOOG"}},
		{"check $GOOG!", []string{"GOOG"}},
		{"$GOOG", []string{"GOOG"}},
		{"$GOOG is up", []string{"GOOG"}},
		{"GOOG", []string{}},
		{"￥GOOG", []string{"GOOG"}},
		{"￥GOOG and $AAPL", []string{"GOOG", "AAPL"}},
	}
	for _, c := range cases {
		if sym := getSymbol(c.text); !sliceEqual(sym,c.symbol) {
			t.Errorf("Expected %q from [%s], got %q\n", c.symbol, c.text, sym)
		}
	}
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}

