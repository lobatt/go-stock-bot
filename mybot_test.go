package main

import (
	"testing"
	"regexp"
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

func TestGetQuote(t *testing.T) {
	cases := []struct {
		input []string
		outputLines int
	} {
		{[]string{"GOOG","AAPL"}, 2},
	}

	pattern := regexp.MustCompile("\\(([[:alnum:]]+)\\)")
	for _, c := range cases {
		msg := getQuote(c.input)
		matches := pattern.FindAllStringSubmatch(msg, -1)
		if len(matches) != c.outputLines {
			t.Errorf("Expected %d result(s) from %q, got %d\n", c.outputLines, c.input, len(matches))
		}
	}
}

