// A small compiler that converts a limited subset of Lisp into C syntax
// e.g. (add 1 2) -> add(1,2)
// Inspired by the super tiny compiler project
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the tiny compiler!")
	tokens, err := Tokenizer("(add 1 2)")
	if err != nil {
		panic(err)
	}
	fmt.Println(tokens)
	ast := Parser(tokens)
	fmt.Println(ast)
}
