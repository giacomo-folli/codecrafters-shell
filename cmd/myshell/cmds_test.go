package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
)

// var argsMap = make(map[string]string)
var cases = []struct {
	in  string
	out string
}{
	{in: `'shell hello'`, out: `shell hello`},
	{in: `"quz  hello"  "bar"`, out: `quz  hello bar`},
	{in: `"quz"  "hello""bar"`, out: `quz hellobar`},
	{in: `"bar"  "shell's"  "foo"`, out: `bar shell's foo`},
	{in: `world\ \ \ \ \ \ script`, out: `world      script`},
	{in: `"before\   after"`, out: `before\   after`},
	{in: `'example\"testhello\"shell'`, out: `example\"testhello\"shell`},
	{in: `'shell\\\nscript'`, out: `shell\\\nscript`},
	{in: `"hello'script'\\n'world"`, out: `hello'script'\n'world`},
	{in: `"hello\"insidequotes"script\"`, out: `hello"insidequotesscript"`},
}

func TestCommandParsing(t *testing.T) {
	for i, ccase := range cases {
		tokenized := _parseArgs(ccase.in)

		got := strings.Join(tokenized, " ")
		want := ccase.out

		if got != want {
			fmt.Println("--------------------------------------------------")
			fmt.Printf("Failed test on: %q, index: %d\n", ccase.in, i)
			fmt.Println(diff.LineDiff(got, want))
			fmt.Println("--------------------------------------------------")
			t.Errorf("\ngot  %q \nwant %q", got, want)
		}
	}
}
