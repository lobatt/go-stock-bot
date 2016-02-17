package main

import (
	"testing"
)

func TestGetSymbol(t *testing.T) {
	cases := []struct{
		text, symbol string
	} {
		{"check $CMCSA", "CMCSA"},
		{"check $CMCSA!", "CMCSA"},
		{"$CMCSA", "CMCSA"},
		{"$CMCSA is up", "CMCSA"},
		{"CMCSA", ""},
		{"￥CMCSA", "CMCSA"},
	}
	for _, c := range cases {
		if sym := getSymbol(c.text); sym != c.symbol {
			t.Errorf("Expected %s from [%s], got %s\n", c.symbol, c.text, sym)
		}
	}
}

