package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		in   []Token
		want ASTNode
	}{
		{
			in: []Token{
				{"paren", "("},
				{"name", "add"},
				{"number", "1"},
				{"number", "2"},
				{"paren", ")"},
			},
			want: ASTNode{
				nodeType: "Program",
				value:    nil,
				params:   nil,
				body: []ASTNode{
					{
						nodeType: "CallExpression",
						value:    "add",
						params: []ASTNode{
							{
								nodeType: "NumberLiteral",
								value:    "1",
								params:   nil,
								body:     nil,
							},
							{
								nodeType: "NumberLiteral",
								value:    "2",
								params:   nil,
								body:     nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		testName := fmt.Sprintf("%s", tC.in)
		t.Run(testName, func(t *testing.T) {
			out := Parser(tC.in)
			if !reflect.DeepEqual(out, tC.want) {
				t.Errorf("got %s, want %s", out, tC.want)
			}
		})
	}
}
