package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		in   string
		want []Token
	}{
		{
			in: "(add 1 2)",
			want: []Token{
				{"paren", "("},
				{"name", "add"},
				{"number", "1"},
				{"number", "2"},
				{"paren", ")"},
			},
		},
	}
	for _, tC := range testCases {
		testName := fmt.Sprintf("%s", tC.in)
		t.Run(testName, func(t *testing.T) {
			out, _ := Tokenizer(tC.in)
			if !reflect.DeepEqual(out, tC.want) {
				t.Errorf("got %s, want %s", out, tC.want)
			}
		})
	}
}
