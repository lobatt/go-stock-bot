/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, _ := slackConnect(os.Args[1])
	fmt.Println("mybot ready, ^C exits")

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Println(err)
			continue
		}

		// see if we're mentioned
		if m.Type == "message" {
			if symbol := getSymbol(m.Text); len(symbol) > 0 {
				m.Text = getQuote(symbol)
				postMessage(ws, m)
			}
		}
	}
}

func getSymbol(msg string) string {
	re := regexp.MustCompile("[\\$|￥]([[:alnum:]]*)")
	matches := re.FindStringSubmatch(msg)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""

}


// Get the quote via Yahoo. You should replace this method to something
// relevant to your team!
func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1p2&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) >= 4 {
		return fmt.Sprintf("%s (%s) is trading at $%s (change %s)", rows[0][0], rows[0][1], rows[0][2], rows[0][3])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}
